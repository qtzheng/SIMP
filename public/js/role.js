mini.parse();
var formRole = new mini.Form("formRole");
var winRole = mini.get("winRole");
var treeRole = mini.get("treeRole");
var treeModule=mini.get("treeModule");
var mainTab = mini.get("tabMain");
var selectRole = undefined;
var selectModule = undefined;

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
    var url = "/System/GetRoleInfo?id=" + selectRole.RoleID;
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
                treeRole.updateNode(selectRole, newNode);
            }
        }
    });
}

function onRoleSelect(e) {
    if (e.node.RoleCode == "Defult") {
        selectRole = undefined;
        return;
    }
    selectRole = e.node;
    var tab = mainTab.getActiveTab();
    if (tab.title == "用户") {
        var iframe = mainTab.getTabIFrameEl(tab);
        iframe.contentWindow.LoadData();
    } else {
        LoadRolePermi();
    }
}

function LoadRolePermi() {
    if (!selectRole)
        return;
    var url = '/System/RoleModule?roleID=' + selectRole.RoleID;
    Ajax({
        url: url,
        success: function(ret) {
            if (ret.Result != 0) {
                mini.alert("模块权限加载出错");
                return;
            }
            if (selectModule)
                OnFunctionLoad();
            CheckModules(ret.Message);
            CloseWaite();
        },
        error: function(q, m, e) {
            CloseWaite();
            mini.alert(m);
        }
    });
}

function onModuleSelect(e) {
    selectModule = e.node;
    gridFunc.load({
        moduleID: selectModule.ID
    });
}

function onModuleCheck(e) {
    if (!selectRole) {
        mini.alert("添加模块前,请选择角色！");
        e.cancel = true;
        return;
    }
    var node = e.node;
    if (e.checked) {
        if (node.PermissionId) {
            if (parentNode.Code = !"system") {
                parentModuleID = parentNode.ID;
            }
            var msg = DeleteRolePer(node.PermissionId);
            if (msg.Result == 0) {
                e.node.PermissionId = undefined;
                ShowTips('删除成功', 500);
            } else {
                ShowTips('删除失败', 500);
                e.cancel = true;
            }
        } else {
            e.cancel = true;
            ShowTips('删除失败', 500);
        }
    } else {
        var data = {
            RoleID: selectRole.RoleID,
            ModuleID: selectModule.ID,
            RoleCode: selectRole.RoleCode
        };
        $.ajax({
            url: '/System/RolePerModuleAdd',
            type: 'post',
            data: data,
            async: false,
            success: function(msg) {
                CloseWaite();
                if (msg.Result == 0) {
                    ShowTips("添加成功");
                    var list = msg.Message;
                    for (var i = list.length - 1; i >= 0; i--) {
                        var per = list[i];
                        var node = treeModule.getNode(per.ModuleID);
                        node.PermissionId = per.PermissionId;
                    };
                } else {
                    e.cancel = true;
                    mini.alert("添加失败");
                }
            },
            error: function(q, m, e) {
                CloseWaite();
                mini.alert(m);
            }
        });
    }
}

function CreateCheck(e) {
    var check = '<input id="ck' + e.record.id + '" type="checkbox" onclick="OnFunctionCheck(\'' + e.record.id + '\',this)"/>';
    var hidPerId = '<input id="hid' + e.record.id + '" type="hidden" value="" />';
    var div = '<div>' + check + hidPerId + '</div>';
    return div;
}

function OnFuncCkClicked(funcId, e) {
    if (!selectRole) {
        mini.alert("请选择角色");
        e.checked = !e.checked;
        return;
    }
    if (e.checked) {
        var msg = SaveRolePer(selectRole.ID, selectModule.ID, funcId, 2);
        if (msg.Result == 0) {
            $('#hid' + funcId).val(msg.Message);
            ShowTips('添加成功', 500);
            AddResourceBtn(funcId, msg.Message);
        } else {
            e.checked = !e.checked;
            ShowTips('添加失败', 500);
            $('#hid' + funcId).val('');
        }
    } else {
        var hidPer = $('#hid' + funcId);
        var perId = hidPer.val();
        if (perId != '') {
            var msg = DeleteRolePer(perId)
            if (msg.Result == 0) {
                hidPer.val('');
                ShowTips('删除成功', 500);
            } else {
                e.checked = !e.checked;
                ShowTips('删除失败', 500);
            }
        }
    }
}

function CheckModules(data) {
    try {
        var nodes = new Array();
        for (var i = data.length - 1; i >= 0; i--) {
           var node=treeModule.getNode(data[i].ModuleID);
            nodes.push(node);
        };
        treeModule.uncheckAllNodes();
        treeModule.checkNodes(nodes);
    } catch (e) {
        alert(e.message);
    }
}

function CheckFuncs(data) {
    var roles = gridFunc.findRows(function(row) {
        RemoveResourceBtn(row.id);
        var hidPer = $('#hid' + row.id);
        var ckFunc = $('#ck' + row.id);
        ckFunc.attr("checked", false);
        hidPer.val('');
        for (var i = 0; i < data.length; i++) {
            if (data[i].funcId == row.id) {
                row.perId = data[i].perId;
                AddResourceBtn(row.id, data[i].perId);
                ckFunc.attr("checked", true);
                hidPer.val(data[i].perId);
                return true;
            }
        }
    });
}

function SaveRolePer(roleId, moduleID, funcId, type) {
    OpenWaite();
    var retData = {
        Result: 1
    };
    var data = {
        RoleID: roleId,
        FunctionID: funcId,
        IsModule: type,
        ModuleID: moduleID
    };
    if (type == 1) {
        data.parentModuleID = treeModule.getParentNode(selectModule).ID;
    }
    $.ajax({
        url: '/System/RolePerAdd',
        type: 'post',
        async: false,
        data: data,
        success: function(ret) {
            retData = ret;
        },
        error: function(q, m, e) {
            CloseWaite();
            mini.alert(m);
        }
    });
    return retData;
}

function DeleteRolePer(perId) {
    OpenWaite();
    var retData;
    $.ajax({
        url: '/System/RolePerDelete?id=' + perId,
        type: 'post',
        async: false,
        success: function(ret) {
            retData = ret;
        },
        error: function(q, m, e) {
            CloseWaite();
            mini.alert(m);
        }
    });
    return retData;
}

function OnFunctionLoad() {
    $.ajax({
        url: '/System/RoleFunc',
        data: {
            roleId: selectRole.ID,
            moduleId: selectModule.ID
        },
        success: function(ret) {
            if (ret.Result == 0)
                CheckFuncs(ret);
            else
                mini.alert("功能点权限加载出错");
        },
        error: function(q, m, e) {
            CloseWaite();
            mini.alert(m);
        }
    });
}