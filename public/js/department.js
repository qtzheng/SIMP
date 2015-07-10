mini.parse();
var formDep = new mini.Form("formDep");
var winDep = mini.get("viewDep");
var treeDep = mini.get("treeDep");
var selectDep = undefined;
function OpenDepAdd() {
    if (!selectDep) {
        mini.alert("请选择上级部门！");
    }
    OpenAddForm(formDep, winDep, "添加部门", "icon-add", function () {
        var parDepId = selectDep.id;
        var parDepName = selectDep.text;
        mini.get("txtParentId").setValue(parDepId);
        $('#spanParentName').html(parDepName);
    });
}
function onDepSelect(e){
	selectDep=e.node;
}
function AddDep () {
    if (!CheckForm(formDep))
        return;
    var data = mini.encode(formDep.getData());
    Ajax({
        url: "/Department/AddDep",
        type: "post",
        data: { role: data },
        success: function (msg) {
            mini.alert("");
        }
    });
}