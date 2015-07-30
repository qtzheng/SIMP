mini.parse();
var treeModule = mini.get("treeModule");
var winModule = mini.get("winModule");
var formModule = new mini.Form("formModule");
var selectModule = undefined;

var gridFunction = mini.get("gridFunction");
var formFunction = new mini.Form("formFunction");
var winFunc = mini.get("winFunc");
var selectFunction = undefined;

function onModuleSelect(e) {
	selectModule = e.node;
	selectFunction = undefined;
	SelectFunc();
}

function OpenModuleAdd() {
	if (!selectModule) {
		mini.alert("请选择父模块");
		return;
	}
	OpenAddForm(formModule, winModule, "添加模块", "icon-add", function() {
		var parId = selectModule.ID;
		var parName = selectModule.Name;
		mini.get("hidPID").setValue(parId);
		mini.get("btnModuleAdd").show();
		mini.get("btnModuleEdit").hide();
		$('#spParentName').html(parName);
		SetAsInput("txtCode");
	});
}

function OpenModuleEdit() {
	if (!selectModule) {
		mini.alert("请选择模块");
		return;
	}
	var url = "/System/ModuleInfo?id=" + selectModule.ID;
	OpenEditForm(url, formModule, winModule, "编辑模块", "icon-edit", function() {
		var Name = treeModule.getParentNode(selectModule).Name;
		$('#spParentName').html(Name);
		mini.get("btnModuleAdd").hide();
		mini.get("btnModuleEdit").show();
		SetAsLabel("txtCode");
	});
}

function AddModule() {
	if (!CheckForm(formModule))
		return;
	var data = formModule.getData();
	Ajax({
		url: "/System/ModuleInsert",
		type: "post",
		data: data,
		success: function(msg) {
			if (msg.Result == 0) {
				ShowTips("保存成功");
				HideWin(winModule);
				var newNode = {
					ID: msg.Message,
					Name: data.Name,
					ParentID: data.ParentID
				};

				treeModule.addNode(newNode, 0, selectModule);
			}
		}
	});
}

function EditModule() {
		if (!CheckForm(formModule))
			return;
		var data = formModule.getData();
		Ajax({
			url: "/System/ModuleUpdate",
			type: "post",
			data: data,
			success: function(msg) {
				if (msg.Result == 0) {
					ShowTips("保存成功");
					HideWin(winModule);
					var newNode = {
						RoleName: data.RoleName,
					};
					treeModule.updateNode(selectModule, newNode);
				}
			}
		});
	}
	//=======================================================================
function OnFuncSelect(e) {
	selectFunction = e.record;
}

function SelectFunc() {
	var moduleId = selectModule.ID;
	if (moduleId == "Defult")
		return;
	gridFunction.load({
		moduleID: moduleId
	});
}

function OpenFunctionAdd() {
	if (!selectModule) {
		mini.alert("请选择所属模块");
		return;
	}
	OpenAddForm(formFunction, winFunc, "添加模块", "icon-add", function() {
		var moduleId = selectModule.ID;
		var moduleName = selectModule.Name;
		mini.get("hidFunModuleID").setValue(moduleId);
		mini.get("btnFunctionAdd").show();
		mini.get("btnFunctionEdit").hide();
		$('#spModuleName').html(moduleName);
		SetAsInput("txtCode");
	});
}

function OpenFunctionEdit() {
	if (!selectFunction) {
		mini.alert("请选择功能点");
		return;
	}
	var url = "/System/FuncInfo?id=" + selectFunction.ID;
	OpenEditForm(url, formFunction, winFunc, "编辑模块", "icon-edit", function() {
		var Name = treeModule.getParentNode(selectModule).Name;
		$('#spModuleName').html(Name);
		mini.get("btnFunctionAdd").hide();
		mini.get("btnFunctionEdit").show();
		SetAsLabel("txtCode");
	});
}

function AddFunction() {
	if (!CheckForm(formFunction))
		return;
	var data = formFunction.getData();
	Ajax({
		url: "/System/FuncInsert",
		type: "post",
		data: data,
		success: function(msg) {
			if (msg.Result == 0) {
				ShowTips("保存成功");
				HideWin(winFunc);
				data.ID = msg.Message;
				gridFunction.addRow(data);
			}
		}
	});
}

function EditFunction() {
	if (!CheckForm(formFunction))
		return;
	var data = formFunction.getData();
	Ajax({
		url: "/System/FuncUpdate",
		type: "post",
		data: data,
		success: function(msg) {
			if (msg.Result == 0) {
				ShowTips("保存成功");
				HideWin(winFunc);
				gridFunction.updateRow(selectFunction, data);
			}
		}
	});
}