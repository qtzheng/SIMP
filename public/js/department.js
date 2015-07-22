mini.parse();
var formDep = new mini.Form("formDep");
var winDep = mini.get("viewDep");
var treeDep = mini.get("treeDep");
var selectDep = undefined;

function OpenDepAdd() {
    if (!selectDep) {
        mini.alert("请选择上级部门！");
    }
    OpenAddForm(formDep, winDep, "添加部门", "icon-add", function() {
        var parDepId = selectDep.ID;
        var parDepName = selectDep.Name;
        mini.get("txtParentId").setValue(parDepId);
        $('#spanParentName').html(parDepName);
        mini.get("depAdd").hide();
        mini.get("depSave").show();
        SetAsInput("txtDepCode");
    });
}

function OpenDepEdit() {
    if (!selectDep) {
        mini.alert("请选择编辑部门！");
    }
    var url = "/System/DepInfo?id=" + selectDep.ID;
    OpenEditForm(url, formDep, winDep, "编辑部门", "icon-edit", function() {
        var depName = treeDep.getParentNode(selectDep).Name;
        $('#spanParentName').html(depName);
        mini.get("depAdd").hide();
        mini.get("depSave").show();
        SetAsLabel("txtDepCode");
    });
}

function onDepSelect(e) {
    selectDep = e.node;
}

function AddDep() {
    if (!CheckForm(formDep))
        return;
    var data = mini.encode(formDep.getData());
    Ajax({
        url: "/System/DepInsert",
        type: "post",
        data: data,
        success: function(msg) {
            if (msg.Result == 0) {
                ShowTips("保存成功");
                HideWin(winDep);
                var newNode = {
                    ID: msg.Message,
                    Name: data.RoleName,
                    ParentID: data.ParentID
                };
                treeDep.addNode(newNode, 0, selectDep);
            }
        }
    });
}

function EditDep() {
    if (!CheckForm(formDep))
        return;
    var data = formDep.getData();
    Ajax({
        url: "/System/DepEdit",
        type: "post",
        data: data,
        success: function(msg) {
            if (msg.Result == 0) {
                ShowTips("保存成功");
                HideWin(winDep);
                var newNode = {
                    Name: data.Name,
                };
                treeDep.updateNode(selectDep, newNode);
            }
        }
    });
}