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
        var parRoleId = selectRole.id;
        var parRoleName = selectRole.text;
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
    var url = "/System/GetRoleInfo";
    OpenEditForm(url, formRole, winRole, "编辑角色", "icon-edit", function() {
        var RoleName = treeRole.getParentNode(selectRole).text;
        $('#spanParentName').html(RoleName);
        mini.get("roleAdd").hide();
        mini.get("roleSave").show();
        SetAsLabel("txtRoleCode");
    })
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
                    id: msg.Message,
                    text: data.RoleName,
                    pid: data.ParentID
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
                    text: data.RoleName,
                };
                treeRole.updateNode(selectRole,newNode);
            }
        }
    });
}

function onRoleSelect(e) {
    selectRole = e.node;
}