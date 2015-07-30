mini.parse();
var formRole = new mini.Form("formRole");
var winRole = mini.get("winRole");
var treeRole = mini.get("treeRole");
var selectRole = undefined;

function OpenRoleAdd() {
    if (!selectRole) {
        mini.alert("请选择父角色");
        return;
    }
    OpenAddForm(formRole, winRole, "添加角色", "icon-add", function() {
        var parRoleId = selectRole.RoleID;
        var parRoleName = selectRole.RoleName;
        mini.get("txtParentId").setValue(parRoleId);
        mini.get("roleAdd").show();
        mini.get("roleSave").hide();
        $('#spanParentName').html(parRoleName);
        SetAsInput("txtRoleCode");
    });
}

function OpenRoleEdit() {
    if (!selectRole) {
        mini.alert("请选择角色");
        return;
    }
    var url = "/System/GetRoleInfo?id="+selectRole.RoleID;
    OpenEditForm(url, formRole, winRole, "编辑角色", "icon-edit", function() {
        var RoleName = treeRole.getParentNode(selectRole).RoleName;
        $('#spanParentName').html(RoleName);
        mini.get("roleAdd").hide();
        mini.get("roleSave").show();
        SetAsLabel("txtRoleCode");
    });
}

function DeleteRole() {

}

function AddRole() {
    if (!CheckForm(formRole))
        return;
    var data = formRole.getData();
    Ajax({
        url: "/System/AddRole",
        type: "post",
        data: data,
        success: function(msg) {
            if (msg.Result == 0) {
                ShowTips("保存成功");
                HideWin(winRole);
                var newNode = {
                    RoleID: msg.Message,
                    RoleName: data.RoleName,
                    ParentID: data.ParentID
                };
                treeRole.addNode(newNode, 0, selectRole);
            }
        }
    });
}

function EditRole() {
    if (!CheckForm(formRole))
        return;
    var data = formRole.getData();
    Ajax({
        url: "/System/EditRole",
        type: "post",
        data: data,
        success: function(msg) {
            if (msg.Result == 0) {
                ShowTips("保存成功");
                HideWin(winRole);
                var newNode = {
                    RoleName: data.RoleName,
                };
                treeRole.updateNode(selectRole,newNode);
            }
        }
    });
}

function onRoleSelect(e) {
    selectRole = e.node;
}
function onModuleSelect(e) {
    selectModule = e.node;
    gridFunc.load({ moduleID: selectModule.ID });
}
function onModuleCheck(e) {
    if (!selectRole) {
        mini.alert("选择模块前,请选择角色！");
        e.cancel = true;
        return;
    }
    var node = e.node;
    if (e.checked) {
        if (node.PermissionId) {
            var msg = DeleteRolePer(node.PermissionId);
            if (msg.Result == 0) {
                e.node.PermissionId = undefined;
                ShowTips('删除成功', 500);
            }
            else {
                ShowTips('删除失败', 500);
                e.cancel = true;
            }
        }
        else {
            e.cancel = true;
            ShowTips('删除失败', 500);
        }
    }
    else {
        var msg = SaveRolePer(selectRole.ID, node.ID, 1);
        if (msg.Result == 0) {
            e.node.PermissionId = msg.Message;
            ShowTips('添加成功', 500);
        }
        else {
            e.cancel = true;
            ShowTips('添加失败', 500);
        }
    }
}