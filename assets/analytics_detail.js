// Load the Visualization API and the piechart package.
google.load('visualization', '1.0', {
    'packages': ['table', 'corechart']
});

// Set a callback to run when the Google Visualization API is loaded.

google.setOnLoadCallback(drawChart);
var romid = "";
var current_version = "";
var chart_drawed = 0;
	
function initCountry(obj) {
    $("#select_country").append($("<option></option>").val("all").html("All Countries"));

    for (var i = 0; i < obj.length; i++) {
        $("#select_country").append($("<option></option>").val(obj[i].RomId).html(obj[i].Country));
        if(romid == obj[i].RomId){
            document.getElementById('select_country').options.selectedIndex = i + 1;
            //$("#select_country").options.selectedIndex = i;
        }
    }
    $("#select_country").append($("<option></option>").val("tmpCuref").html("tmpCuref"));
    $('#select_country').on('change', function(e) {
        romid = this.value;
        location.href="/analytics_viewer_details?serial_selected=" +romid + "&version_selected=" + current_version;
    });
}

function initVersion() {
    
    $("#select_version").append($("<option></option>").val(4).html("2015"));
    $("#select_version").append($("<option></option>").val(3).html("2014"));
	$("#select_version").append($("<option></option>").val(5).html("2016"));
	    var selection = 0;
    if (current_version == 4) {
        selection = 0;
    } else if (current_version == 3) {
        selection = 1;
    } else if (current_version == 5) {
        selection = 2;
    }
    document.getElementById('select_version').options.selectedIndex = selection;
	console.log(selection);
    $('#select_version').on('change', function(e) {
        current_version = this.value;
        location.href="/analytics_viewer_details?serial_selected=" +"all" + "&version_selected=" + current_version;
    });
}

    // Callback that creates and populates a data table, 
// instantiates the pie chart, passes in the data and
// draws it.
function drawChart() {
    var selection = current_version;
    if (selection == "") {
        selection = "4";
    }
    $.ajax({
        url: "/analytics_api_get_device/version/" + selection,
        dataType: "json",
        success: function(data) {
            initCountry(data);
        }
    });
    initVersion();
    drawOTA();
   /* drawVersion();
    drawProfilesByAge();
    drawNbChildProfiles();
   */
}

function drawOTA() {

    var dataOTA = new google.visualization.DataTable();
    dataOTA.addColumn('string', 'OTA version');
    dataOTA.addColumn('number', 'Count');

    for (var i = 0; i < this.otaRepartition; i++) {
//	{{range $idx,$val := .otaRepartition}}
    dataOTA.addRow([this.otaRepartition[i].name, parseInt(this.otaRepartition[i].Count)]);
   // {{end}}

	}

    var optionsOTA = {
        'title': 'OTA Repartition'
    };
    // Instantiate and draw our chart, passing in some options.
    var chartOTA = new google.visualization.Table(document.getElementById('OTA'));
    chartOTA.draw(dataOTA, optionsOTA);

}
/*
function drawVersion() {

    var dataVersion = new google.visualization.DataTable();
    dataVersion.addColumn('string', 'Kurio system version');
    dataVersion.addColumn('number', 'Count');

    {{range $idx,$val := .systemVersion}}
    dataVersion.addRow(['{{$val.Name}}', parseInt('{{$val.Count}}')]);
    {{end}}

    var options = {
        'title': 'OTA Repartition'
    };

    // Instantiate and draw our chart, passing in some options.
    var chart = new google.visualization.Table(document.getElementById('version'));
    chart.draw(dataVersion, options);

}

function drawProfilesByAge() {

    var data = new google.visualization.DataTable();
    data.addColumn('string', 'Profiles by Age');
    data.addColumn('number', 'Nb Profiles');

    {{range $idx,$val := .profilesByAge}}
    data.addRow(['{{$val.Name}}', parseInt('{{$val.Count}}')]);
    {{end}}

    var options = {
        'title': 'Profiles by Age'
    };

    // Instantiate and draw our chart, passing in some options.
    var chart = new google.visualization.LineChart(document.getElementById('profilesByAge'));
    chart.draw(data, options);

}



function drawNbChildProfiles() {

    var data = new google.visualization.DataTable();
    data.addColumn('string', 'Number of childs');
    data.addColumn('number', 'Count (Deleted profile included)');

    {{range $idx,$val := .nbChildProfiles}}
    data.addRow(['{{$val.Name}}', parseInt('{{$val.Count}}')]);
    {{end}}

    var options = {
        'title': 'Number of childs'
    };

    // Instantiate and draw our chart, passing in some options.
    var chart = new google.visualization.Table(document.getElementById('nbChildProfiles'));
    chart.draw(data, options);

}

function getAppByAge() {
    
    var ageElement = document.getElementById("ageSelected");
    var ageSelected = ageElement.options[ageElement.selectedIndex].value;

    var genderElement = document.getElementById("gender");
    var genderSelected = genderElement.options[genderElement.selectedIndex].value;

    $.ajax({
        url: "/analytics_api_get_app_by_age_sex?serial_selected="+romid + "&age_selected="+ageSelected+"&gender="+genderSelected,
        dataType: "json",
        success: function(data) {
            drawAppByAge(data);
        }
    });
}

function drawAppByAge(obj) {
    //apps bestUninstalled
    var data = new google.visualization.DataTable();
    data.addColumn('string', 'Package Name');
    data.addColumn('number', 'Time Spent (in min)');
    data.addColumn('number', 'Launch count');

    for (var i = 0; i < obj.length; i++) {
        var timeSpentMinute = parseInt(obj[i].TotalTimeSpent / 60);
        data.addRow([obj[i].PackageName, timeSpentMinute, parseInt(obj[i].TotalLaunchCount)]);
    }

    // // Instantiate and draw our chart, passing in some options.
    var chart = new google.visualization.BarChart(document.getElementById('appsByAge'));
    chart.draw(data);
}
 */
 
