const ctx = document.getElementById('myChart');

new Chart(ctx, {
    type: 'line',
    data: {
        labels: ['Jan', 'Fev', 'Mar', 'Avr', 'Mai', 'Jun'],
        datasets: [{
            label: 'Ventes',
            data: [12, 19, 3, 5, 2, 10],
            borderColor: 'rgba(255, 255, 255, 1)',
            borderWidth: 2,
            fill: true,
            tension: 0.3,

            backgroundColor: (context) => {
                const chart = context.chart;
                const {
                    ctx,
                    chartArea
                } = chart;

                if (!chartArea) return null;

                const gradient = ctx.createLinearGradient(0, chartArea.top, 0, chartArea.bottom);

                gradient.addColorStop(0, 'rgba(255, 255, 255, 0.6)'); // près de la courbe
                gradient.addColorStop(1, 'rgba(255,255,255,0)');       // vers le bas

                return gradient;
            }
        }]
    },

    options: {
        responsive: true,
        plugins: {
            legend: { display: false }
        },
        scales: {
            x: { display: false },
            y: { display: false }
        }
    }
});

