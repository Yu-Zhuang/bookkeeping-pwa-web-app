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
            data = myJson.data
            for(let i = 0; i < data.line.length; i++) {
                lineLabel.push(data.line[i].month)
                lineData.push(data.line[i].total)
            }
            for(let i = 0; i < data.pie.length; i++) {
                pieLabel.push(data.pie[i].class)
                pieData.push(data.pie[i].total)
            }
    }
    else {
        alert("無法載入資料")
        return
    }
    // line chart
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
            backgroundColor: [
            'rgb(255, 99, 132)',
            'rgb(54, 162, 235)',
            'rgb(255, 205, 86)'
            ],
            hoverOffset: 4
        }]
        };

    const config2 = {
        type: 'doughnut',
        data: data2,
    };
    var myChart2 = new Chart(document.getElementById('pieChart'), config2)  
}