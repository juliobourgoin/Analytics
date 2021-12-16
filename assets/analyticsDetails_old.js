// Load the Visualization API and the piechart package.
google.load('visualization', '1.0', {
    'packages': ['table', 'corechart']
});

// Set a callback to run when the Google Visualization API is loaded.

google.setOnLoadCallback(drawChart);

function initCountry(obj) {
    $("#select_country").append($("<option></option>").val("").html("All Country"));

    for (var i = 0; i < obj.length; i++) {
        $("#select_country").append($("<option></option>").val(obj[i].RomId).html(obj[i].Country));
        if(romid == obj[i].RomId){
            document.getElementById('select_country').options.selectedIndex = i + 1;
            //$("#select_country").options.selectedIndex = i;
        }
    }
    $("#select_country").append($("<option></option>").val("tmpCuref").html("Curef temp"));
    $('#select_country').on('change', function(e) {
        romid = this.value;
        location.href="/analytics_viewer_details?serial_selected=" +romid;
    });
}
    // Callback that creates and populates a data table, 
// instantiates the pie chart, passes in the data and
// draws it.
function drawChart() {

    $.ajax({
        url: "/analytics_api_get_device",
        dataType: "json",
        success: function(data) {
            initCountry(data);
        }
    });
    drawOTA();
}

function drawOTA() {

    var dataOTA = new google.visualization.DataTable();
    dataOTA.addColumn('string', 'OTA version');
    dataOTA.addColumn('number', 'Count');

    {{range $idx,$val := .otaRepartition}}
    dataOTA.addRow(['{{$val.Name}}', parseInt('{{$val.Count}}')]);
    {{end}}

    var optionsProfile = {
        'title': 'OTA Repartition'
    };

    // Instantiate and draw our chart, passing in some options.
    var chartProfile = new google.visualization.Table(document.getElementById('OTA'));
    chartProfile.draw(dataProfile, optionsProfile);

}