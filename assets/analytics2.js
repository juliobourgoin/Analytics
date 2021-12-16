// Load the Visualization API and the piechart package.
google.load('visualization', '1.0', {
	'packages': ['table', 'corechart']
});

var romid = ""
	var current_version = ""

	var load_1 = false
	var load_2 = false
	var load_3 = false
	var load_4 = false
	// Specific variable to update analytics_viewer page one time
	var chart_drawed = 0

	processUser();

function processUser() {
	window.history.forward();
}

// Set a callback to run when the Google Visualization API is loaded.
google.setOnLoadCallback(drawChart);

function drawRomId(obj) {

	var data = new google.visualization.DataTable();
	var cssClassNames = {
		'headerRow': 'italic-darkblue-font large-font bold-font',
		'tableRow': '',
		'oddTableRow': 'beige-background',
		'selectedTableRow': 'orange-background large-font',
		'hoverTableRow': '',
		'headerCell': 'gold-border',
		'tableCell': '',
		'rowNumberCell': 'underline-blue-font'
	};
	if (document.title != "Extract Data") {
		data.addColumn('string', 'Model');
		data.addColumn('number', 'Count');
		for (var i = 0; i < obj.length; i++) {
			if (obj[i].Country != "%)%") {
				if (obj[i].Country != "999") {
					if (obj[i].Country != "710") {
						if (obj[i].Country != "888") {
							if (obj[i].Country != "101") {
								if (obj[i].Country != "102") {
									if (obj[i].Country != "13") {
										if (obj[i].Country != "221") {
												data.addRow([obj[i].Country, obj[i].Count]);
												//console.log(obj[i].Count);
												nblig = data.getNumberOfRows() - 1;
												//	console.log("le nb lig est" +nblig);
												nbcol = data.getNumberOfColumns() - 1;
												//	console.log("le nb col est" +nbcol);
												if (current_version == 2) {
													data.setCell(nblig, nbcol - 1, String([obj[i].Country]), String([obj[i].Country]), {
														'className': 'bold-white-font kurio-background'
													});
													//data.setCell(nblig, nbcol,0, obj[i].Count, {'className': 'bold-font'});
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

		// Set chart options
		var options = {
			'title': 'Activation of Device By Country',
			'width': 900,
			'height': 700,
			'is3D': true
		};

		// Instantiate and draw our chart, passing in some options.

		chart_div = document.getElementById('tabRomId');
			var chart = new google.visualization.PieChart(document.getElementById('tabRomId'));
			chart.draw(data, options);
			google.visualization.events.addListener(chart, 'select', function () {
				selected_pie = data.getValue(chart.getSelection()[0].row, 0);
				romid = selected_pie.substring(selected_pie.length - 3, selected_pie.length);
				drawChartRomID();

			});
			chart.draw(data, options);
		

	} else {
		data.addColumn('string', 'Country');
		data.addColumn('number', 'Count');
		for (var i = 0; i < obj.length; i++) {
			data.addRow([obj[i].Country, obj[i].Count])
		}

		// Set chart options
		var options = {
			'title': 'By rom ID',
			'width': 600,
			'height': 600
		};

		// Instantiate and draw our chart, passing in some options.

		var chart = new google.visualization.Table(document.getElementById('tabRomId'));
		chart.draw(data, options);

	}
}

function initCountry(obj) {
	$("#select_country").empty();
	$("#select_country").append($("<option></option>").val("all").html("All Countries"));

	for (var i = 0; i < obj.length; i++) {
		if(obj[i].Country != ""){
			$("#select_country").append($("<option></option>").val(obj[i].RomId).html(obj[i].Country));
		}
	}

	$("#select_country").append($("<option></option>").val("tmpCuref").html("tmpCuref"));
	
	$('#select_country').on('change', function (e) {
		romid = this.value;
		drawChartRomID();
	});

}

function initVersion() {

	var selection = 0;
	if (current_version == 4) {
		selection = 0;
	} else if (current_version == 3) {
		selection = 1;
	} else if (current_version == 5) {
		selection = 2;
	} else if (current_version == 2) {
		selection = 3;
	} else if (current_version == 6) {
		selection = 4;
	}

	document.getElementById('select_version').options.selectedIndex = selection;

	$('#select_version').on('change', function (e) {
		current_version = this.value;
		romid = "all";
		chart_drawed = 1;
		drawChart();
	});
}

function drawProfileInfo(obj) {

	var dataProfile = new google.visualization.DataTable();
	dataProfile.addColumn('string', 'Subject / all include deleted profiles');
	dataProfile.addColumn('number', 'Count');

	dataProfile.addRow(["Total", obj.Total]);
	dataProfile.addRow(["Boy", obj.Total - obj.Girl]);
	dataProfile.addRow(["Girl", obj.Girl]);
	dataProfile.addRow(["TimeControlEnabled", obj.TimeControlEnabled]);
	dataProfile.addRow(["TimeSlotEnabled", obj.TimeSlotEnabled]);
	dataProfile.addRow(["AdvancedSettingsEnabled", obj.AdvancedSettingsEnabled]);
	dataProfile.addRow(["AllowUsbStatus", obj.AllowUsbStatus]);
	dataProfile.addRow(["AuthorizeAdsStatus", obj.AuthorizeAdsStatus]);
	dataProfile.addRow(["AutoAuthorizeStatus", obj.AutoAuthorizeStatus]);
	dataProfile.addRow(["ProfileChanged", obj.ProfileChanged]);
	dataProfile.addRow(["ChildProtectedByPWD", obj.ChildProtectedByPWD]);
	dataProfile.addRow(["WebListActivated", obj.WebListActivated]);
	dataProfile.addRow(["Deleted", obj.Deleted]);

	var optionsProfile = {
		'title': 'All Profile infos'
	};

	// Instantiate and draw our chart, passing in some options.
	var chartProfile = new google.visualization.Table(document.getElementById('profile'));
	chartProfile.draw(dataProfile, optionsProfile);
}

function drawActivation(obj) {
	var dataActivation = new google.visualization.DataTable();
	dataActivation.addColumn('string', 'ActivationDate');
	dataActivation.addColumn('number', 'Total Count');
	dataActivation.addColumn('number', 'Daily Count');
	var total = 0;

	for (var i = 0; i < obj.length; i++) {
		total = total + obj[i].Count;
		dataActivation.addRow([obj[i].ActivationDate, total, parseInt(obj[i].Count)]);
	}

	// Set chart options
	var optionsActivation = {
		'title': 'Activation : you can zoom with mouse wheel to see details',
		'height': 500,
		explorer: {
			maxZoomOut: 2,
			keepInBounds: true,
			curveType: 'function'
		}
	};

	// Instantiate and draw our chart, passing in some options.
	var chartActivation = new google.visualization.LineChart(document.getElementById('tabActivation'));
	chartActivation.draw(dataActivation, optionsActivation);

	load_1 = true;

	if (load_1 && load_2 && load_3 && load_4) {
		console.log("hidden Activation");
		document.getElementById('loading_icon').style.visibility = "hidden";
	}

}

function drawActivityStatus(obj) {
	//activity status
	/*  var dataAppPref = new google.visualization.DataTable();
	var dataAppPref2 = new google.visualization.DataTable();
	var cssClassNames = {
	'headerRow': 'italic-darkblue-font large-font bold-font',
	'tableRow': '',
	'oddTableRow': 'beige-background',
	'selectedTableRow': 'orange-background large-font',
	'hoverTableRow': '',
	'headerCell': 'gold-border',
	'tableCell': '',
	'rowNumberCell': 'underline-blue-font'};
	dataAppPref.addColumn('string', 'Package Name');
	dataAppPref.addColumn('number', 'Time in Minute');
	dataAppPref.addColumn('number', 'Launch Count');

	dataAppPref2.addColumn('string', 'Package Name');
	dataAppPref2.addColumn('number', 'Time in Minute');
	dataAppPref2.addColumn('number', 'Launch Count');


	for (var i = 0; i < obj.length; i++) {
	var timeSpent = parseInt(obj[i].TotalTimeSpent);
	dataAppPref.addRow([obj[i].PackageName, parseInt(timeSpent / 60), parseInt(obj[i].TotalLaunchCount)]);
	dataAppPref2.addRow([obj[i].PackageName, parseInt(timeSpent / 60), parseInt(obj[i].TotalLaunchCount)]);

	nblig=dataAppPref.getNumberOfRows()-1;
	nbcol=dataAppPref.getNumberOfColumns()-1;

	dataAppPref.setCell(nblig, nbcol-2,String([obj[i].PackageName]),String([obj[i].PackageName]), {'className': 'bold-white-font kurio-background'});
	dataAppPref.setCell(nblig, nbcol-1,0, parseInt(timeSpent / 60), {'className': 'bold-font'});
	dataAppPref.setCell(nblig, nbcol,0, parseInt(obj[i].TotalLaunchCount), {'className': 'bold-font'});

	}

	var view = new google.visualization.DataView(dataAppPref);
	view.setColumns([0, 1, 2]);

	var table = new google.visualization.Table(document.getElementById('appPrefTable'));
	table.draw(view);

	// Set chart options
	var optionsAppPref = {
	'title': '50 First ActivityStatus order by time spent',
	'height': 1000
	};
	var options = {
	sort: 'event'
	};
	// Instantiate and draw our chart, passing in some options.
	var chartAppPref = new google.visualization.BarChart(document.getElementById('appPref'));
	chartAppPref.draw(dataAppPref2, optionsAppPref);

	google.visualization.events.addListener(table, 'sort',

	function(event) {
	dataAppPref.sort([{
	column: event.column,
	desc: !event.ascending
	}]);
	chartAppPref.draw(view);
	});
	google.visualization.events.addListener(table, 'select', selectHandler);

	load_2 = true;

	if (load_1 && load_2 && load_3 && load_4) {
	console.log("hidden activityStatus");
	document.getElementById('loading_icon').style.visibility="hidden";
	}
	 */
	//activity status
	var dataAppPref = new google.visualization.DataTable();
	var cssClassNames = {
		'headerRow': 'italic-darkblue-font large-font bold-font',
		'tableRow': '',
		'oddTableRow': 'beige-background',
		'selectedTableRow': 'orange-background large-font',
		'hoverTableRow': '',
		'headerCell': 'gold-border',
		'tableCell': '',
		'rowNumberCell': 'underline-blue-font'
	};

	dataAppPref.addColumn('string', 'Package Name');
	dataAppPref.addColumn('number', 'Time in Minute');
	dataAppPref.addColumn('number', 'Launch Count');

	for (var i = 0; i < obj.length; i++) {
		var timeSpent = parseInt(obj[i].TotalTimeSpent);
		dataAppPref.addRow([obj[i].PackageName, parseInt(timeSpent / 60), parseInt(obj[i].TotalLaunchCount)]);
		nblig = dataAppPref.getNumberOfRows() - 1;
		nbcol = dataAppPref.getNumberOfColumns() - 1;

		dataAppPref.setCell(nblig, nbcol - 2, String([obj[i].PackageName]), String([obj[i].PackageName]), {
			'className': 'bold-white-font kurio-background'
		});
		//dataAppPref.setCell(nblig, nbcol-1,0, parseInt(timeSpent / 60), {'className': 'bold-font'});
		//dataAppPref.setCell(nblig, nbcol,0, parseInt(obj[i].TotalLaunchCount), {'className': 'bold-font'});
	}

	var view = new google.visualization.DataView(dataAppPref);
	view.setColumns([0, 1, 2]);

	var table = new google.visualization.Table(document.getElementById('appPrefTable'));
	table.draw(view);

	// Set chart options
	var optionsAppPref = {
		'title': '50 First ActivityStatus order by time spent',
		'width': 1500,
		'height': 1000
	};

	// Instantiate and draw our chart, passing in some options.
	var chartAppPref = new google.visualization.BarChart(document.getElementById('appPref'));
	chartAppPref.draw(dataAppPref, optionsAppPref);

	google.visualization.events.addListener(table, 'sort',

		function (event) {
		dataAppPref.sort([{
					column: event.column,
					desc: !event.descending
				}
			]);
		chartAppPref.draw(view);
	});

	load_2 = true;

	if (load_1 && load_2 && load_3 && load_4) {
		console.log("hidden activityStatus");
		document.getElementById('loading_icon').style.visibility = "hidden";
	}
}
function selectHandler(e) {
	alert('A table row was selected');
}


// Callback that creates and populates a data table,
// instantiates the pie chart, passes in the data and
// draws it.
function drawChart() {

	var selection = current_version;
	if (selection == "") {
		selection = "2";
	}

	$.ajax({
		url: "/analytics_api_get_device/version/" + selection,
		dataType: "json",
		success: function (data) {
			drawRomId(data);
			if (chart_drawed == 0) {
				initVersion();
			}
			initCountry(data);

		}
	});

	/*$.ajax({
	url: "/analytics_api_get_profiles/version/" + selection,
	dataType: "json",
	success: function(data) {
	drawProfileInfo(data);
	}
	});*/

	drawChartRomID();
	myval = document.getElementById("serial1");
	myval.value = document.getElementById("version_selected");
}
function createElement(str) {
	var div = document.createElement('div');
	div.innerHTML = str;
	var container = document.createDocumentFragment();
	for (var i = 0; i < div.childNodes.length; i++) {
		var node = div.childNodes[i].cloneNode(true);
		container.appendChild(node);
	}
	return container.childNodes;
}

function drawChartRomID() {

	load_1 = false;
	load_2 = false;
	load_3 = true;
	load_4 = true;
	document.getElementById('loading_icon').style.visibility = 'visible';
	var selectedVersion = current_version;
	if (selectedVersion == "") {
		selectedVersion = "2";
	}

	var selectedRomid = romid;
	if (selectedRomid == "") {
		selectedRomid = "all";
	}

	myval = document.getElementById("serial1");
	myval.value = selectedRomid;
	myval = document.getElementById("version1");
	myval.value = selectedVersion;
	console.log("/version/" + selectedVersion + "/serial/" + selectedRomid);
	
		var div = document.getElementById("profile");
		var divtab = document.getElementById("tableaux");
		if (div) {}
		else {
			divtab.innerHTML += '<div id="profile"></div>';
			divtab.innerHTML += '<div id="appPrefTable"></div>';
			divtab.innerHTML += '<div id="appPref"></div>';
			document.getElementById('profile').style.visibility = 'visible';
			document.getElementById('appPrefTable').style.visibility = 'visible';
			document.getElementById('appPref').style.visibility = 'visible';
		}
		$.ajax({

			url: "/analytics_api_get_activation/version/" + selectedVersion + "/serial/" + selectedRomid,
			dataType: "json",
			success: function (data) {
				drawActivation(data);
			}
		});

		$.ajax({
			url: "/analytics_api_get_activity/version/" + selectedVersion + "/serial/" + selectedRomid,
			dataType: "json",
			success: function (data) {
				drawActivityStatus(data);
			}
		});

	

}
