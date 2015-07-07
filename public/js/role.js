mini.parse();
var formRole = new mini.Form("formRole");
var winRole = mini.get("winRole");
var treeRole = mini.get("treeRole");
var selectRole = undefined;
function OpenRoleAdd() {
    if (!selectRole) {
        mini.alert("请选择父角色");
    }
    OpenAddForm(formRole, winRole, "添加角色", "icon-add", function () {
        var parRoleId = selectRole.id;
        var parRoleName = selectRole.text;
        mini.get("txtParentId").setValue(parRoleId);
        $('#spanParentName').html(parRoleName);
    });
}
function onRoleSelect(e) {
    selectRole = e.node;
}