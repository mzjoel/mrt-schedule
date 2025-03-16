document.addEventListener("DOMContentLoaded", () => {
    fetchStations();
    document.getElementById("stations").addEventListener("change", fetchSchedules);
});

async function fetchStations() {
    try {
        const response = await fetch("http://localhost:8080/v1/api/stations");
        const result = await response.json();
        
        if (Array.isArray(result.Data)) {
            const stationSelect = document.getElementById("stations");
            stationSelect.innerHTML = `<option value="">-- Pilih Stasiun --</option>`;

            result.Data.forEach(station => {
                const option = document.createElement("option");
                option.value = station.nid;
                option.textContent = station.title;
                stationSelect.appendChild(option);
            });
        }
    } catch (error) {
        console.error("Error fetching stations:", error);
    }
}

async function fetchSchedules() {
    const stationId = document.getElementById("stations").value;
    if (!stationId) return;

    try {
        const response = await fetch(`http://localhost:8080/v1/api/stations/${stationId}`);
        const result = await response.json();
        
        if (Array.isArray(result.Data)) {
            displaySchedules(result.Data);
        }
    } catch (error) {
        console.error("Error fetching schedules:", error);
    }
}

function displaySchedules(schedules) {
    const container = document.getElementById("schedule-container");
    container.innerHTML = ""; // Bersihkan sebelum menampilkan yang baru

    if (schedules.length === 0) {
        container.innerHTML = "<p>Tidak ada jadwal tersedia.</p>";
        return;
    }

    schedules.forEach(schedule => {
        const scheduleItem = document.createElement("div");
        scheduleItem.classList.add("schedule-item");
        scheduleItem.textContent = `${schedule.time}`;
        container.appendChild(scheduleItem);
    });
}