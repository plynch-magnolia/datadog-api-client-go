package test

import (
	"context"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
)

func testingRoleCreateAttributes(ctx context.Context, t *testing.T) *datadog.RoleCreateAttributes {
	rca := datadog.NewRoleCreateAttributesWithDefaults()
	rca.SetName(*tests.UniqueEntityName(ctx, t))
	return rca
}

func deleteRole(ctx context.Context, t *testing.T, roleID string) {
	_, err := Client(ctx).RolesApi.DeleteRole(ctx, roleID)
	if err == nil {
		return
	}
	t.Logf("Error disabling Role: %v, Another test may have already deleted this role: %s", roleID, err.Error())
}

func TestRoleLifecycle(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// first, test creating a role
	rca := testingRoleCreateAttributes(ctx, t)
	rcd := datadog.NewRoleCreateDataWithDefaults()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreateRequestWithDefaults()
	rcp.SetData(*rcd)
	rr, httpresp, err := Client(ctx).RolesApi.CreateRole(ctx, *rcp)
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(ctx, t, rid)

	rrAttributes := rrData.GetAttributes()
	assert.Equal(rca.GetName(), rrAttributes.GetName())

	// now, test updating it
	updatedRoleName := rca.GetName() + "-updated"
	rua := datadog.NewRoleUpdateAttributesWithDefaults()
	rua.SetName(updatedRoleName)
	rud := datadog.NewRoleUpdateDataWithDefaults()
	rud.SetAttributes(*rua)
	rud.SetId(rid)
	rup := datadog.NewRoleUpdateRequestWithDefaults()
	rup.SetData(*rud)

	urr, httpresp, err := Client(ctx).RolesApi.UpdateRole(ctx, rid, *rup)
	if err != nil {
		t.Fatalf("Error updating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	urData := urr.GetData()
	urAttributes := urData.GetAttributes()
	assert.Equal(rua.GetName(), urAttributes.GetName())

	// now, test getting it
	grr, httpresp, err := Client(ctx).RolesApi.GetRole(ctx, rid)
	if err != nil {
		t.Fatalf("Error getting Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	getData := grr.GetData()
	getAttributes := getData.GetAttributes()
	assert.Equal(updatedRoleName, getAttributes.GetName())

	// now, test filtering for it in the list call
	rsr, httpresp, err := Client(ctx).RolesApi.ListRoles(ctx, *datadog.NewListRolesOptionalParameters().
		WithFilter(updatedRoleName).
		WithPageSize(1).
		WithPageNumber(0).
		WithSort(datadog.ROLESSORT_MODIFIED_AT_DESCENDING))
	if err != nil {
		t.Fatalf("Error listing roles %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.NotEmpty(rsr.GetData())
	listAttributes := rsr.GetData()[0].GetAttributes()
	assert.Equal(rua.GetName(), listAttributes.GetName())
	rrMetaAttributes := rsr.GetMeta()
	page := rrMetaAttributes.GetPage()
	assert.GreaterOrEqual(page.GetTotalCount(), int64(1))
	assert.GreaterOrEqual(page.GetTotalFilteredCount(), int64(1))

	// now, test deleting it
	// no response payload
	httpresp, err = Client(ctx).RolesApi.DeleteRole(ctx, rid)
	if err != nil {
		t.Fatalf("Error deleting Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(204, httpresp.StatusCode)
}

func TestRolePermissionsLifecycle(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// first, create a role
	rca := testingRoleCreateAttributes(ctx, t)
	rcd := datadog.NewRoleCreateDataWithDefaults()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreateRequestWithDefaults()
	rcp.SetData(*rcd)
	rr, httpresp, err := Client(ctx).RolesApi.CreateRole(ctx, *rcp)
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(ctx, t, rid)

	// find a permission
	permissions, httpresp, err := Client(ctx).RolesApi.ListPermissions(ctx)
	if err != nil {
		t.Fatalf("Error listing permissions: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(permissions.HasData())
	assert.NotEmpty(permissions.GetData())

	permission := permissions.GetData()[0]
	pid := permission.GetId()

	// add a permission to the role
	rtp := datadog.NewRelationshipToPermissionWithDefaults()
	rtpd := datadog.NewRelationshipToPermissionDataWithDefaults()
	rtpd.SetId(pid)
	rtp.SetData(*rtpd)

	crrtps, httpresp, err := Client(ctx).RolesApi.AddPermissionToRole(ctx, rid, *rtp)
	if err != nil {
		t.Fatalf("Error creating permission relation: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Contains(crrtps.GetData(), permission)

	// get all permissions for the role
	lrrtps, httpresp, err := Client(ctx).RolesApi.ListRolePermissions(ctx, rid)
	if err != nil {
		t.Fatalf("Error listing permission relations: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.Contains(lrrtps.GetData(), permission)

	// remove the permission from the role
	drrtps, httpresp, err := Client(ctx).RolesApi.RemovePermissionFromRole(ctx, rid, *rtp)
	if err != nil {
		t.Fatalf("Error remove permission relation: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.NotContains(drrtps.GetData(), permission)
}

func TestRoleUsersLifecycle(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// first, create a role
	rca := testingRoleCreateAttributes(ctx, t)
	rcd := datadog.NewRoleCreateDataWithDefaults()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreateRequestWithDefaults()
	rcp.SetData(*rcd)
	rr, httpresp, err := Client(ctx).RolesApi.CreateRole(ctx, *rcp)
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(ctx, t, rid)

	// create a user
	uca := testingUserCreateAttributes(ctx, t)
	ucd := datadog.NewUserCreateDataWithDefaults()
	ucd.SetAttributes(*uca)
	ucr := datadog.NewUserCreateRequestWithDefaults()
	ucr.SetData(*ucd)
	ur, httpresp, err := Client(ctx).UsersApi.CreateUser(ctx, *ucr)
	if err != nil {
		assert.IsType(datadog.GenericOpenAPIError{}, err, "%v", err)
		t.Fatalf("Error creating User %s: Response %s: %v", uca.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(httpresp.StatusCode, 201)
	urData := ur.GetData()
	uid := urData.GetId()
	defer disableUser(ctx, t, uid)

	// add a user to the role
	rtu := datadog.NewRelationshipToUserWithDefaults()
	rtud := datadog.NewRelationshipToUserDataWithDefaults()
	rtud.SetId(uid)
	rtu.SetData(*rtud)

	crrtus, httpresp, err := Client(ctx).RolesApi.AddUserToRole(ctx, rid, *rtu)
	if err != nil {
		assert.IsType(datadog.GenericOpenAPIError{}, err, "%v", err)
		t.Fatalf("Error creating user relation: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	found := false
	for _, ua := range crrtus.GetData() {
		if uid == ua.GetId() {
			found = true
			break
		}
	}
	assert.True(found, "user %s not found in relation to role %s", uid, rid)

	// get all users for the role
	lrrtus, httpresp, err := Client(ctx).RolesApi.ListRoleUsers(ctx, rid)
	if err != nil {
		t.Fatalf("Error listing permission relations: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)

	found = false
	for _, ua := range lrrtus.GetData() {
		if uid == ua.GetId() {
			found = true
			break
		}
	}
	assert.True(found, "user %s not found in relation to role %s", uid, rid)

	// remove the permission from the role
	drrtus, httpresp, err := Client(ctx).RolesApi.RemoveUserFromRole(ctx, rid, *rtu)
	if err != nil {
		assert.IsType(datadog.GenericOpenAPIError{}, err, "%v", err)
		t.Fatalf("Error remove permission relation: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	for _, ua := range drrtus.GetData() {
		if uid == ua.GetId() {
			t.Fatalf("User %s must not exist in the relation %s", uid, rid)
		}
	}
}

func TestListRolesErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
	}{
		"403 Forbidden": {WithFakeAuth, 403},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).RolesApi.ListRoles(ctx)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestCreateRoleErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()

	// first, test creating a role
	rca := testingRoleCreateAttributes(ctx, t)
	rcd := datadog.NewRoleCreateDataWithDefaults()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreateRequestWithDefaults()
	rcp.SetData(*rcd)

	// invalid role without data
	rcp400 := &datadog.RoleCreateRequest{}

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
		Body               *datadog.RoleCreateRequest
	}{
		"400 Bad Request": {WithTestAuth, 400, rcp400},
		"403 Forbidden":   {WithFakeAuth, 403, rcp},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			rr, httpresp, err := Client(ctx).RolesApi.CreateRole(ctx, *tc.Body)
			// make sure that we clean everything on error
			if 200 == httpresp.StatusCode {
				rrData := rr.GetData()
				rid := rrData.GetId()
				defer deleteRole(ctx, t, rid)
			}
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestGetRoleErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// valid role ID
	rca := testingRoleCreateAttributes(ctx, t)
	rcd := datadog.NewRoleCreateDataWithDefaults()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreateRequestWithDefaults()
	rcp.SetData(*rcd)
	rr, httpresp, err := Client(ctx).RolesApi.CreateRole(ctx, *rcp)
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(ctx, t, rid)

	// bad role ID
	rid404 := "00000000-dead-beef-dead-ffffffffffff"
	_, httpresp, err = Client(ctx).RolesApi.GetRole(ctx, rid404)
	assert.Equal(404, httpresp.StatusCode)

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
		RoleID             string
	}{
		"403 Forbidden": {WithFakeAuth, 403, rid},
		"404 Not found": {WithTestAuth, 404, rid404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).RolesApi.GetRole(ctx, tc.RoleID)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestUpdateRoleErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// valid role ID
	rca := testingRoleCreateAttributes(ctx, t)
	rcd := datadog.NewRoleCreateDataWithDefaults()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreateRequestWithDefaults()
	rcp.SetData(*rcd)
	rr, httpresp, err := Client(ctx).RolesApi.CreateRole(ctx, *rcp)
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(ctx, t, rid)

	// working update payload
	updatedRoleName := rca.GetName() + "-updated"
	rua := datadog.NewRoleUpdateAttributesWithDefaults()
	rua.SetName(updatedRoleName)
	rud := datadog.NewRoleUpdateDataWithDefaults()
	rud.SetAttributes(*rua)
	rud.SetId(rid)
	rup := datadog.NewRoleUpdateRequestWithDefaults()
	rup.SetData(*rud)

	// bad role ID
	rid404 := "00000000-dead-beef-dead-ffffffffffff"
	_, httpresp, err = Client(ctx).RolesApi.GetRole(ctx, rid404)
	assert.Equal(404, httpresp.StatusCode)
	rud404 := datadog.NewRoleUpdateDataWithDefaults()
	rud404.SetAttributes(*rua)
	rud404.SetId(rid404)
	rup404 := datadog.NewRoleUpdateRequestWithDefaults()
	rup404.SetData(*rud404)

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
		RoleID             string
		Body               *datadog.RoleUpdateRequest
	}{
		"400 Bad Request":            {WithTestAuth, 400, rid, datadog.NewRoleUpdateRequestWithDefaults()},
		"403 Forbidden":              {WithFakeAuth, 403, rid, rup},
		"404 Bad Role ID in Path":    {WithTestAuth, 404, rid404, rup404},
		"422 Bad Role ID in Request": {WithTestAuth, 422, rid, rup404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).RolesApi.UpdateRole(ctx, tc.RoleID, *tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestDeleteRoleErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// valid role ID
	rca := testingRoleCreateAttributes(ctx, t)
	rcd := datadog.NewRoleCreateDataWithDefaults()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreateRequestWithDefaults()
	rcp.SetData(*rcd)
	rr, httpresp, err := Client(ctx).RolesApi.CreateRole(ctx, *rcp)
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(ctx, t, rid)

	// bad role ID
	rid404 := "00000000-dead-beef-dead-ffffffffffff"
	_, httpresp, err = Client(ctx).RolesApi.GetRole(ctx, rid404)
	assert.Equal(404, httpresp.StatusCode)

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
		RoleID             string
	}{
		"403 Forbidden": {WithFakeAuth, 403, rid},
		"404 Not found": {WithTestAuth, 404, rid404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			httpresp, err := Client(ctx).RolesApi.DeleteRole(ctx, tc.RoleID)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestListRolePermissionsErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// valid role ID
	rca := testingRoleCreateAttributes(ctx, t)
	rcd := datadog.NewRoleCreateDataWithDefaults()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreateRequestWithDefaults()
	rcp.SetData(*rcd)
	rr, httpresp, err := Client(ctx).RolesApi.CreateRole(ctx, *rcp)
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(ctx, t, rid)

	// bad role ID
	rid404 := "00000000-dead-beef-dead-ffffffffffff"
	_, httpresp, err = Client(ctx).RolesApi.GetRole(ctx, rid404)
	assert.Equal(404, httpresp.StatusCode)

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
		RoleID             string
	}{
		"403 Forbidden": {WithFakeAuth, 403, rid},
		"404 Not found": {WithTestAuth, 404, rid404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).RolesApi.ListRolePermissions(ctx, tc.RoleID)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestAddPermissionToRoleErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// valid role ID
	rca := testingRoleCreateAttributes(ctx, t)
	rcd := datadog.NewRoleCreateDataWithDefaults()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreateRequestWithDefaults()
	rcp.SetData(*rcd)
	rr, httpresp, err := Client(ctx).RolesApi.CreateRole(ctx, *rcp)
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(ctx, t, rid)

	// bad role ID
	rid404 := "00000000-dead-beef-dead-ffffffffffff"
	_, httpresp, err = Client(ctx).RolesApi.GetRole(ctx, rid404)
	assert.Equal(404, httpresp.StatusCode)

	rtp := datadog.NewRelationshipToPermissionWithDefaults()
	rtpd := datadog.NewRelationshipToPermissionDataWithDefaults()
	rtpd.SetId(rid404)
	rtp.SetData(*rtpd)

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
		Body               *datadog.RelationshipToPermission
	}{
		"400 Bad Request": {WithTestAuth, 400, datadog.NewRelationshipToPermissionWithDefaults()},
		"403 Forbidden":   {WithFakeAuth, 403, rtp},
		"404 Not found":   {WithTestAuth, 404, rtp},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).RolesApi.AddPermissionToRole(ctx, rid, *tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestRemovePermissionFromRoleErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// valid role ID
	rca := testingRoleCreateAttributes(ctx, t)
	rcd := datadog.NewRoleCreateDataWithDefaults()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreateRequestWithDefaults()
	rcp.SetData(*rcd)
	rr, httpresp, err := Client(ctx).RolesApi.CreateRole(ctx, *rcp)
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(ctx, t, rid)

	// bad role ID
	rid404 := "00000000-dead-beef-dead-ffffffffffff"
	_, httpresp, err = Client(ctx).RolesApi.GetRole(ctx, rid404)
	assert.Equal(404, httpresp.StatusCode)

	// find a permission
	permissions, httpresp, err := Client(ctx).RolesApi.ListPermissions(ctx)
	if err != nil {
		t.Fatalf("Error listing permissions: Response %s: %v", err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	assert.True(permissions.HasData())
	assert.NotEmpty(permissions.GetData())

	permission := permissions.GetData()[0]
	pid := permission.GetId()

	rtp := datadog.NewRelationshipToPermissionWithDefaults()
	rtpd := datadog.NewRelationshipToPermissionDataWithDefaults()
	rtpd.SetId(pid)
	rtp.SetData(*rtpd)

	// invalid permission data
	rtp400 := datadog.NewRelationshipToPermissionWithDefaults()
	rtpd400 := datadog.NewRelationshipToPermissionDataWithDefaults()
	rtpd400.SetId("11111111-dead-beef-dead-ffffffffffff")
	rtp400.SetData(*rtpd400)

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
		RoleID             string
		Body               *datadog.RelationshipToPermission
	}{
		"400 Bad Request":        {WithTestAuth, 400, rid, datadog.NewRelationshipToPermissionWithDefaults()},
		"400 Invalid Permission": {WithTestAuth, 400, rid, rtp400},
		"403 Forbidden":          {WithFakeAuth, 403, rid, rtp},
		"404 Role Not Found":     {WithTestAuth, 404, rid404, rtp},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).RolesApi.RemovePermissionFromRole(ctx, tc.RoleID, *tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestListRoleUsersErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// valid role ID
	rca := testingRoleCreateAttributes(ctx, t)
	rcd := datadog.NewRoleCreateDataWithDefaults()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreateRequestWithDefaults()
	rcp.SetData(*rcd)
	rr, httpresp, err := Client(ctx).RolesApi.CreateRole(ctx, *rcp)
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(ctx, t, rid)

	// bad role ID
	rid404 := "00000000-dead-beef-dead-ffffffffffff"
	_, httpresp, err = Client(ctx).RolesApi.GetRole(ctx, rid404)
	assert.Equal(404, httpresp.StatusCode)

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
		RoleID             string
	}{
		"403 Forbidden": {WithFakeAuth, 403, rid},
		"404 Not found": {WithTestAuth, 404, rid404},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).RolesApi.ListRoleUsers(ctx, tc.RoleID)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestAddUserToRoleErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// valid role ID
	rca := testingRoleCreateAttributes(ctx, t)
	rcd := datadog.NewRoleCreateDataWithDefaults()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreateRequestWithDefaults()
	rcp.SetData(*rcd)
	rr, httpresp, err := Client(ctx).RolesApi.CreateRole(ctx, *rcp)
	if err != nil {
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(ctx, t, rid)

	// bad role ID
	rid404 := "00000000-dead-beef-dead-ffffffffffff"
	_, httpresp, err = Client(ctx).RolesApi.GetRole(ctx, rid404)
	assert.Equal(404, httpresp.StatusCode)

	rtu := datadog.NewRelationshipToUserWithDefaults()
	rtud := datadog.NewRelationshipToUserDataWithDefaults()
	rtud.SetId(rid404)
	rtu.SetData(*rtud)

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
		Body               *datadog.RelationshipToUser
	}{
		"400 Bad Request": {WithTestAuth, 400, &datadog.RelationshipToUser{}},
		"403 Forbidden":   {WithFakeAuth, 403, rtu},
		"404 Not found":   {WithTestAuth, 404, rtu},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).RolesApi.AddUserToRole(ctx, rid, *tc.Body)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}

func TestRemoveUserFromRoleErrors(t *testing.T) {
	ctx, finish := tests.WithTestSpan(context.Background(), t)
	defer finish()
	ctx, finish = WithRecorder(WithTestAuth(ctx), t)
	defer finish()
	assert := tests.Assert(ctx, t)

	// valid role ID
	rca := testingRoleCreateAttributes(ctx, t)
	rcd := datadog.NewRoleCreateDataWithDefaults()
	rcd.SetAttributes(*rca)
	rcp := datadog.NewRoleCreateRequestWithDefaults()
	rcp.SetData(*rcd)
	rr, httpresp, err := Client(ctx).RolesApi.CreateRole(ctx, *rcp)
	if err != nil {
		assert.IsType(datadog.GenericOpenAPIError{}, err, "%v", err)
		t.Fatalf("Error creating Role %s: Response %s: %v", rca.GetName(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(200, httpresp.StatusCode)
	rrData := rr.GetData()
	rid := rrData.GetId()
	defer deleteRole(ctx, t, rid)

	// bad role ID
	rid404 := "00000000-dead-beef-dead-ffffffffffff"
	_, httpresp, err = Client(ctx).RolesApi.GetRole(ctx, rid404)
	assert.Equal(404, httpresp.StatusCode)

	// create a user
	uca := testingUserCreateAttributes(ctx, t)
	ucd := datadog.NewUserCreateDataWithDefaults()
	ucd.SetAttributes(*uca)
	ucr := datadog.NewUserCreateRequestWithDefaults()
	ucr.SetData(*ucd)
	ur, httpresp, err := Client(ctx).UsersApi.CreateUser(ctx, *ucr)
	if err != nil {
		assert.IsType(datadog.GenericOpenAPIError{}, err, "%v", err)
		t.Fatalf("Error creating User %s: Response %s: %v", uca.GetEmail(), err.(datadog.GenericOpenAPIError).Body(), err)
	}
	assert.Equal(httpresp.StatusCode, 201)
	urData := ur.GetData()
	uid := urData.GetId()
	defer disableUser(ctx, t, uid)

	rtu := datadog.NewRelationshipToUserWithDefaults()
	rtud := datadog.NewRelationshipToUserDataWithDefaults()
	rtud.SetId(uid)
	rtu.SetData(*rtud)

	testCases := map[string]struct {
		Ctx                func(context.Context) context.Context
		ExpectedStatusCode int
		RoleID             string
		Body               *datadog.RelationshipToUser
	}{
		"400 Bad Request": {WithTestAuth, 400, rid, &datadog.RelationshipToUser{}},
		"403 Forbidden":   {WithFakeAuth, 403, rid, rtu},
		"404 Bad role":    {WithTestAuth, 404, rid404, rtu},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx, finish := WithRecorder(tc.Ctx(ctx), t)
			defer finish()
			assert := tests.Assert(ctx, t)

			_, httpresp, err := Client(ctx).RolesApi.RemoveUserFromRole(ctx, tc.RoleID, *tc.Body)
			assert.IsType(datadog.GenericOpenAPIError{}, err, "%v", err)
			assert.Equal(tc.ExpectedStatusCode, httpresp.StatusCode)
			apiError, ok := err.(datadog.GenericOpenAPIError).Model().(datadog.APIErrorResponse)
			assert.True(ok)
			assert.NotEmpty(apiError.GetErrors())
		})
	}
}
