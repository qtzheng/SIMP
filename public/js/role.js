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
    OpenAddForm(formRole, winRole, "添加角色", "icon-add", function () {
        var parRoleId = selectRole.id;
        var parRoleName = selectRole.text;
        mini.get("txtParentId").setValue(parRoleId);
        $('#spanParentName').html(parRoleName);
    });
}
function OpenRoleEdit() {
    if (!selectRole) {
        mini.alert("请选择角色");
        return;
    }
    var url = "/Role/GetRoleInfo";
    OpenEditForm(url, formRole, winRole, "编辑角色", "icon-edit", function () {
        var RoleName = treeRole.getParentNode(selectRole).text;
        $('#spanParentName').html(RoleName);
    })
}
function DeleteRole() {

}
function AddRole() {
    if (!CheckForm(formRole))
        return;
    var data = mini.encode(formRole.getData());
    Ajax({
        url: "/Role/AddRole",
        type: "post",
        data: { role: data },
        success: function (msg) {
            mini.alert("");
        }
    });
}
function onRoleSelect(e) {
    selectRole = e.node;
}
