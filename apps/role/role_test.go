package role

import (
	"testing"
)

func TestHasPermission(t *testing.T) {
	set := NewRoleSet()
	r := &Role{
		Spec: &CreateRoleRequest{
			Permissions: []*Permission{
				{
					//Service: "cmdb",
					//Featrues: []*Featrue{
					//	{
					//		Resource: "secret",
					//		Action:   "list",
					//	},
					//	{
					//		Resource: "secret",
					//		Action:   "get",
					//	},
					//	{
					//		Resource: "secret",
					//		Action:   "create",
					//	},
					//},
					Service:  "cmdb",
					AllowAll: true,
				},
			},
		},
	}

	set.Add(r)

	perm, role := set.HasPermission(&PermissionRequest{
		Service:  "cmdb",
		Resource: "secret",
		Action:   "create",
	})

	t.Log(role)

	if perm != true {
		t.Fatal("has perm error")
	}
}
