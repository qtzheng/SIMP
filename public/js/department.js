mini.parse();
var formDep = new mini.Form("formDep");
var winDep = mini.get("viewDep");
var treeDep = mini.get("treeDep");
var selectDep = undefined;
var winUser = mini.get("viewEmp");
var formUser = new mini.Form("formUser");
var gridUser = mini.get("gridUser");

function OpenDepAdd() {
    if (!selectDep) {
        mini.alert("请选择上级部门！");
    }
    OpenAddForm(formDep, winDep, "添加部门", "icon-add", function() {
        var parDepId = selectDep.ID;
        var parDepName = selectDep.Name;
        mini.get("txtParentId").setValue(parDepId);
        $('#spanParentName').html(parDepName);
        mini.get("depAdd").show();
        mini.get("depSave").hide();
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
    var data = formDep.getData();
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
                    Name: data.Name,
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
}OpenUserAdd
//=============================================
function SelectUser() {
    try {
        var key = mini.get("txtKey").getValue();
        var depIds = new Array();
        if (selectDep != undefined&&selectDep.ID != "Defult") {
            depIds.push(selectDep.ID);
            var nodes = treeDep.getAllChildNodes(selectDep);
            nodes.each(function(index, e) {
                depIds.push(e.ID);
            });
        }
        gridUser.load({
            key: key,
            depIds: depIds.join(",")
        });
    } catch (e) {
        mini.alert(e.Message);
    }
}

function OpenUserAdd() {
    OpenAddForm(formUser, winUser, "添加员工", "icon-add", function() {
        mini.get("btnUserAdd").show();
        mini.get("btnUserEdit").hide();
        mini.get("btnSetPas").hide();
        SetAsInput("txtLoginName");
        SetAsInput("txtJobNumber");
    });
}

function OpenUserEdit() {
    var selectRow=gridUser.getSelected();
    if (!selectRow) {
        mini.alert("请选择编辑用户！");
        return;
    }
    var url = "/System/UserInfo?id=" + selectRow.UserID;
    OpenEditForm(url, formUser, winUser, "编辑员工", "icon-edit", function() {
        mini.get("btnUserAdd").hide();
        mini.get("btnUserEdit").show();
        mini.get("btnSetPas").show();
        SetAsLabel("txtLoginName");
        SetAsLabel("txtJobNumber");
    },function(data){
        var ids=data.RoleIDs.join(',');
        mini.get("tsRoleID").setValue(ids);
    });
}

function AddUser() {
    if (!CheckForm(formUser))
        return;
    var data = formUser.getData();
    Ajax({
        url: "/System/UserInsert",
        type: "post",
        data: data,
        success: function(msg) {
            if (msg.Result == 0) {
                ShowTips("保存成功");
                HideWin(winUser);
                gridUser.addRow(data);
            }
        }
    });
}

function EditUser() {
    if (!CheckForm(formUser))
        return;
    var data = formUser.getData();
    Ajax({
        url: "/System/UserUpdate",
        type: "post",
        data: data,
        success: function(msg) {
            if (msg.Result == 0) {
                ShowTips("保存成功");
                HideWin(winUser);
                 var selectRow=gridUser.getSelected();
                gridUser.updateRow(selectRow, data);
            }
        }
    });
}