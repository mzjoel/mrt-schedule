function displayMRTInfo(data) {
    const container = document.getElementById("mrt-info");
    if (!data) {
        container.innerHTML = "<p>Gagal memuat data MRT</p>";
        return;
    }

    const content = `
        <h2>${data.name}</h2>
        <p>Rute: ${data.route}</p>
        <p>Status: ${data.status}</p>
    `;
    container.innerHTML = content;
}
