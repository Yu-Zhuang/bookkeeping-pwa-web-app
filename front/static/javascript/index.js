function getChart() {
    // line chart
    const labels = [
        'January',
        'February',
        'March',
        'April',
        'May',
        'June',
    ];
    const data = {
        labels: labels,
        datasets: [{
            label: '每月花費',
            backgroundColor: 'rgb(255, 99, 132)',
            borderColor: 'rgb(255, 99, 132)',
            data: [3500, 4000, 3300, 6000, 4500, 3700, 3400],
        }]
    };
    const config = {
        type: 'line',
        data,
        options: {}
    };
    var myChart = new Chart(document.getElementById('lineChart'), config)        
    getPieChart()
}

function getPieChart() {
    const data2 = {
        labels: [
            '食',
            '衣',
            '住'
        ],
        datasets: [{
            label: 'My First Dataset',
            data: [300, 50, 100],
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