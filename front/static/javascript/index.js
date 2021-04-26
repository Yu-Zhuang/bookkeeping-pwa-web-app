async function getChart() {
    let lineLabel = []
    let lineData = []
    let pieLabel = []
    let pieData = []
    // get data
    const res = await fetch((hostUrl + 'api/getChartData'), {
        method: 'GET',
    })
    if (res.status == 200) {
            myJson = await res.json()
            line = myJson.line
            pie = myJson.pie
            for(let i = 0; i < line.length; i++) {
                lineLabel.push(line[i].month)
                lineData.push(line[i].total)
            }
            for(let i = 0; i < pie.length; i++) {
                pieLabel.push(pie[i].class)
                pieData.push(pie[i].total)
            }
            getLineChart(lineLabel, lineData, pieLabel, pieData)
    }
    else {
        alert("無法載入資料")
        return
    }
}

function getLineChart(lineLabel, lineData, pieLabel, pieData) {
    const labels = lineLabel;
    const data = {
        labels: labels,
        datasets: [{
            label: '月花費',
            backgroundColor: 'rgb(255, 99, 132)',
            borderColor: 'rgb(255, 99, 132)',
            data: lineData,
        }]
    };
    const config = {
        type: 'line',
        data,
        options: {}
    };
    var myChart = new Chart(document.getElementById('lineChart'), config)        
    getPieChart(pieLabel, pieData)
}

function getPieChart(pieLabels, pieData) {
    const data2 = {
        labels: pieLabels,
        datasets: [{
            label: '分佈圖',
            data: pieData,
            backgroundColor: Object.values(Utils.CHART_COLORS),
            hoverOffset: 6
        }]
        };

    const config2 = {
        type: 'doughnut',
        data: data2,
    };
    var myChart2 = new Chart(document.getElementById('pieChart'), config2)  
}