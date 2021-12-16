google["load"]("visualization", "1.0", {
	"packages" : [ "table", "corechart" ]
});
/** @type {string} */
var romid = "";
/** @type {string} */
var current_version = "";
/** @type {boolean} */
var load_1 = false;
/** @type {boolean} */
var load_2 = false;
/** @type {boolean} */
var load_3 = false;
/** @type {boolean} */
var load_4 = false;
/** @type {number} */
var chart_drawed = 0;
processUser();
/**
 * @return {undefined}
 */
function processUser() {
	window["history"]["forward"]();
}
google["setOnLoadCallback"](drawChart);

/**
 * @param {Object} result
 * @return {undefined}
 */
function drawRomId(result) {
	console.log("drawRomId");
	var r20 = new google["visualization"].DataTable;
	var tableCell = {
		"headerRow" : "italic-darkblue-font large-font bold-font",
		"tableRow" : "",
		"oddTableRow" : "beige-background",
		"selectedTableRow" : "orange-background large-font",
		"hoverTableRow" : "",
		"headerCell" : "gold-border",
		"tableCell" : "",
		"rowNumberCell" : "underline-blue-font"
	};
	if (document["title"] != "Extract Data") {
		r20["addColumn"]("string", "Model");
		r20["addColumn"]("number", "Count");
		/** @type {number} */
		var i = 0;
		for (; i < result["length"]; i++) {
			if (result[i]["Country"] != "%)%") {
				if (result[i]["Country"] != "999") {
					if (result[i]["Country"] != "710") {
						if (result[i]["Country"] != "888") {
							if (result[i]["Country"] != "101") {
								if (result[i]["Country"] != "102") {
									if (result[i]["Country"] != "13") {
										if (result[i]["Country"] != "221") {
											r20["addRow"]([
													result[i]["Country"],
													result[i]["Count"] ]);
											/** @type {number} */
											nblig = r20["getNumberOfRows"]() - 1;
											/** @type {number} */
											nbcol = r20["getNumberOfColumns"]() - 1;
											if (current_version == 2) {
												r20["setCell"]
														(
																nblig,
																nbcol - 1,
																String([ result[i]["Country"] ]),
																String([ result[i]["Country"] ]),
																{
																	"className" : "bold-white-font kurio-background"
																});
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
		var defaults = {
			"title" : "Activation of Device By Country",
			"backgroundColor" : "transparent",
			"height" : 500,
			"is3D" : true
		};
		chart_div = document["getElementById"]("tabRomId");
		var _window = new google["visualization"].PieChart(
				document["getElementById"]("tabRomId"));
		_window["draw"](r20, defaults);
		google["visualization"]["events"]["addListener"]
				(_window, "select",
						function() {
							selected_pie = r20["getValue"](
									_window["getSelection"]()[0]["row"], 0);
							romid = selected_pie["substring"](
									selected_pie["length"] - 3,
									selected_pie["length"]);
							drawChartRomID();
						});
		_window["draw"](r20, defaults);
	} else {
		r20["addColumn"]("string", "Country");
		r20["addColumn"]("number", "Count");
		/** @type {number} */
		i = 0;
		for (; i < result["length"]; i++) {
			r20["addRow"]([ result[i]["Country"], result[i]["Count"] ]);
		}
		defaults = {
			"title" : "By rom ID",
			"height" : 500,
		};
		_window = new google["visualization"].Table(document["getElementById"]
				("tabRomId"));
		_window["draw"](r20, defaults);
	}
}
/**
 * @param {Object} result
 * @return {undefined}
 */
function initCountry(result) {
	console.log("initCountry");
	$("#select_country")["empty"]();
	$("#select_country")["append"]($("<option></option>")["val"]("all")["html"]
			("All Countries"));
	/** @type {number} */
	var i = 0;
	for (; i < result["length"]; i++) {
		if (result[i]["Country"] != "") {
			$("#select_country")["append"]($("<option></option>")["val"]
					(result[i].RomId)["html"](result[i].Country));
		}
	}
	$("#select_country")["append"]
			($("<option></option>")["val"]("tmpCuref")["html"]("tmpCuref"));
	$("#select_country")["on"]("change", function(dataAndEvents) {
		romid = this["value"];
		drawChartRomID();
	});
}
/**
 * @return {undefined}
 */
function initVersion() {
	console.log("initVersion");
	/** @type {number} */
	var bp = 0;
	if (current_version == 4) {
		/** @type {number} */
		bp = 0;
	} else {
		if (current_version == 3) {
			/** @type {number} */
			bp = 1;
		} else {
			if (current_version == 5) {
				/** @type {number} */
				bp = 2;
			} else {
				if (current_version == 2) {
					/** @type {number} */
					bp = 3;
				} else {
					if (current_version == 6) {
						/** @type {number} */
						bp = 4;
					}
				}
			}
		}
	}
	/** @type {number} */
	document["getElementById"]("select_version")["options"]["selectedIndex"] = bp;
	$("#select_version")["on"]("change", function(dataAndEvents) {
		current_version = this["value"];
		/** @type {string} */
		romid = "all";
		/** @type {number} */
		chart_drawed = 1;
		drawChart();

	});
}
/**
 * @param {Object} dataAndEvents
 * @return {undefined}
 */
function drawProfileInfo(dataAndEvents) {
	console.log("drawProfileInfo");
	var resp = new google["visualization"].DataTable;
	resp["addColumn"]("string", "Subject / all include deleted profiles");
	resp["addColumn"]("number", "Count");
	resp["addRow"]([ "Total", dataAndEvents["Total"] ]);
	resp["addRow"]([ "Boy", dataAndEvents["Total"] - dataAndEvents["Girl"] ]);
	resp["addRow"]([ "Girl", dataAndEvents["Girl"] ]);
	resp["addRow"]
			([ "TimeControlEnabled", dataAndEvents["TimeControlEnabled"] ]);
	resp["addRow"]([ "TimeSlotEnabled", dataAndEvents["TimeSlotEnabled"] ]);
	resp["addRow"]([ "AdvancedSettingsEnabled",
			dataAndEvents["AdvancedSettingsEnabled"] ]);
	resp["addRow"]([ "AllowUsbStatus", dataAndEvents["AllowUsbStatus"] ]);
	resp["addRow"]
			([ "AuthorizeAdsStatus", dataAndEvents["AuthorizeAdsStatus"] ]);
	resp["addRow"]([ "AutoAuthorizeStatus",
			dataAndEvents["AutoAuthorizeStatus"] ]);
	resp["addRow"]([ "ProfileChanged", dataAndEvents["ProfileChanged"] ]);
	resp["addRow"]([ "ChildProtectedByPWD",
			dataAndEvents["ChildProtectedByPWD"] ]);
	resp["addRow"]([ "WebListActivated", dataAndEvents["WebListActivated"] ]);
	resp["addRow"]([ "Deleted", dataAndEvents["Deleted"] ]);
	var options = {
		"title" : "All Profile infos"

	};
	var collection = new google["visualization"].Table(
			document["getElementById"]("profile"));
	collection["draw"](resp, options);
}
/**
 * @param {Object} events
 * @return {undefined}
 */
function drawActivation(events) {
	console.log("drawActivation");
	var resp = new google["visualization"].DataTable;
	resp["addColumn"]("string", "ActivationDate");
	resp["addColumn"]("number", "Total Count");
	resp["addColumn"]("number", "Daily Count");
	/** @type {number} */
	var copies = 0;
	/** @type {number} */
	var i = 0;
	for (; i < events["length"]; i++) {
		copies = copies + events[i]["Count"];
		resp["addRow"]([ events[i]["ActivationDate"], copies,
				parseInt(events[i].Count) ]);
	}
	var options = {
		"backgroundColor" : "transparent",
		"title" : "Activation : you can zoom with mouse wheel to see details",
		"height" : 500,
		explorer : {

			maxZoomOut : 2,
			keepInBounds : true,
			curveType : "function"
		}
	};
	var collection = new google["visualization"].LineChart(
			document["getElementById"]("tabActivation"));
	collection["draw"](resp, options);
	/** @type {boolean} */
	load_1 = true;
	if (load_1 && (load_2 && (load_3 && load_4))) {
		console["log"]("hidden Activation");
		/** @type {string} */
		document["getElementById"]("loading_icon")["style"]["visibility"] = "hidden";
	}
}
/**
 * @param {Object} obj
 * @return {undefined}
 */
function drawActivityStatus(obj) {
	console.log("drawActivityStatus");
	var options = new google["visualization"].DataTable;
	var tableCell = {
		"headerRow" : "italic-darkblue-font large-font bold-font",
		"tableRow" : "",
		"oddTableRow" : "beige-background",
		"selectedTableRow" : "orange-background large-font",
		"hoverTableRow" : "",
		"headerCell" : "gold-border",
		"tableCell" : "",
		"rowNumberCell" : "underline-blue-font"
	};
	options["addColumn"]("string", "Package Name");
	options["addColumn"]("number", "Time in Minute");
	options["addColumn"]("number", "Launch Count");
	/** @type {number} */
	var i = 0;
	for (; i < obj["length"]; i++) {
		/** @type {number} */
		var days = parseInt(obj[i].TotalTimeSpent);
		options["addRow"]([ obj[i]["PackageName"], parseInt(days / 60),
				parseInt(obj[i].TotalLaunchCount) ]);
		/** @type {number} */
		nblig = options["getNumberOfRows"]() - 1;
		/** @type {number} */
		nbcol = options["getNumberOfColumns"]() - 1;
		options["setCell"](nblig, nbcol - 2, String([ obj[i]["PackageName"] ]),
				String([ obj[i]["PackageName"] ]), {
					"className" : "bold-white-font kurio-background"
				});
		// options["setCell"](nblig, nbcol - 1, String([obj[i]["timeSpentMinute"]]), String([obj[i]["timeSpentMinute"]]), {
		// "className" : "bold-green-font "
		// });
		options["setCell"](nblig, nbcol - 1, undefined, null, {
			"className" : "normal-font",

		});

	}
	var p = new google["visualization"].DataView(options);
	p["setColumns"]([ 0, 1, 2 ]);
	var walker = new google["visualization"].Table(document["getElementById"]
			("appPrefTable"));
	walker["draw"](p);
	var base_conf = {
		"backgroundColor" : "transparent",

		"title" : "50 First ActivityStatus order by time spent",
		"width" : 1200,
		"height" : 1200,
	};
	var $w = new google["visualization"].BarChart(document["getElementById"]
			("appPref"));
	$w["draw"](options, base_conf);
	google["visualization"]["events"]["addListener"](walker, "sort", function(
			columns) {
		options["sort"]([ {
			column : columns["column"],
			desc : !columns["descending"]
		} ]);
		$w["draw"](p);
	});
	/** @type {boolean} */
	load_2 = true;
	if (load_1 && (load_2 && (load_3 && load_4))) {
		console["log"]("hidden activityStatus");
		/** @type {string} */
		document["getElementById"]("loading_icon")["style"]["visibility"] = "hidden";

	}
}
/**
 * @param {?} event
 * @return {undefined}
 */
function selectHandler(event) {
}
/**
 * @return {undefined}
 */
function drawChart() {
	console.log("drawChart");
	var type = current_version;
	if (type == "") {
		/** @type {string} */
		type = "2";
	}
	$["ajax"]({
		url : "/analytics_api_get_device/version/" + type,
		dataType : "json",
		/**
		 * @param {Object} response
		 * @return {undefined}
		 */
		success : function(response) {
			drawRomId(response);
			if (chart_drawed == 0) {
				initVersion();
			}
			initCountry(response);

		}
	});
	drawChartRomID();
	myval = document["getElementById"]("serial1");
	myval["value"] = document["getElementById"]("version_selected");
}
/**
 * @param {?} val
 * @return {?}
 */
function createElement(val) {
	var res = document["createElement"]("div");
	res["innerHTML"] = val;
	var el = document["createDocumentFragment"]();
	/** @type {number} */
	var i = 0;
	for (; i < res["childNodes"]["length"]; i++) {
		var ig = res["childNodes"][i]["cloneNode"](true);
		el["appendChild"](ig);
	}
	return el["childNodes"];
}
/**
 * @return {undefined}
 */
function drawChartRomID() {
	/** @type {boolean} */
	load_1 = false;
	/** @type {boolean} */
	load_2 = false;
	/** @type {boolean} */
	load_3 = true;
	/** @type {boolean} */
	load_4 = true;
	/** @type {string} */
	document["getElementById"]("loading_icon")["style"]["visibility"] = "visible";
	var url = current_version;
	if (url == "") {
		/** @type {string} */
		url = "2";
	}
	var name = romid;
	if (name == "") {
		/** @type {string} */
		name = "all";
	}
	myval = document["getElementById"]("serial1");
	myval["value"] = name;
	myval = document["getElementById"]("version1");
	myval["value"] = url;
	console["log"]("/version/" + url + "/serial/" + name);
	var profile = document["getElementById"]("profile");
	var PDIV = document["getElementById"]("tableaux");
	if (profile) {
	} else {
		PDIV["innerHTML"] += '<div id="profile"></div>';
		PDIV["innerHTML"] += '<div id="appPrefTable"></div>';
		PDIV["innerHTML"] += '<div id="appPref"></div>';
		/** @type {string} */
		document["getElementById"]("profile")["style"]["visibility"] = "visible";
		/** @type {string} */
		document["getElementById"]("appPrefTable")["style"]["visibility"] = "visible";
		/** @type {string} */
		document["getElementById"]("appPref")["style"]["visibility"] = "visible";
	}
	$["ajax"]({
		url : "/analytics_api_get_activation/version/" + url + "/serial/"
				+ name,
		dataType : "json",
		/**
		 * @param {Object} event
		 * @return {undefined}
		 */
		success : function(event) {
			setTimeout(drawActivation(event), 1000);
		}
	});

	$["ajax"]({
		url : "/analytics_api_get_activity/version/" + url + "/serial/" + name,
		dataType : "json",
		/**
		 * @param {Object} walkers
		 * @return {undefined}
		 */
		success : function(walkers) {
			setTimeout(drawActivityStatus(walkers), 1000);
		}
	});
};