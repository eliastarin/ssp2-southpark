(() => {
  const $ = (id) => document.getElementById(id);
  const statusEl = $("status");
  const logEl = $("log");
  let timer = null;

  const palette = {
    ok: "ok",
    err: "err",
  };
  const log = (msg, type = "ok") => {
  const line = document.createElement("div");
  line.className = `line ${palette[type] || ""}`;

  // Clean up unwanted text
  msg = msg.replace(/\(queued\)/gi, "").trim();

  // Time formatting — only hours and minutes
  const now = new Date();
  const time = now.toLocaleTimeString([], { 
    hour: "2-digit", 
    minute: "2-digit", 
    hour12: true
  });

  // Optional icon styling
  const icon = type === "ok" ? "💬" : type === "err" ? "⚠️" : "📦";

  // Build output — no brackets, no seconds
  line.innerHTML = `
    <span class="timestamp">${time}</span> 
    ${icon} ${msg}
  `;

  logEl.prepend(line);
};


  const setOnline = (online) => {
    statusEl.textContent = online ? "online" : "offline";
    statusEl.style.borderColor = online ? "rgba(34,211,238,.5)" : "rgba(255,255,255,.12)";
    statusEl.style.color = online ? "#22d3ee" : "#b5b7c5";
  };

  async function pingHealth() {
    try {
      const r = await fetch("/health", { cache: "no-store" });
      setOnline(r.ok);
    } catch {
      setOnline(false);
    }
  }

  async function send(author, body) {
    try {
      const res = await fetch("/messages", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ author, body }),
      });
      const data = await res.json().catch(() => ({}));
      if (!res.ok) throw new Error(`${res.status} ${res.statusText}`);
      log(`${author}: "${body}" (${data.status || "queued"})`, "ok");
    } catch (e) {
      log(`ERROR → ${e.message}`, "err");
    }
  }

  $("sendBtn").onclick = () => {
    const author = $("author").value;
    const body = $("body").value || "Respect my authoritah!";
    send(author, body);
  };

  $("autoBtn").onclick = () => {
    if (timer) return;
    $("autoBtn").disabled = true;
    $("stopBtn").disabled = false;

    const quotes = [
      ["Cartman", "Respect my authoritah!"],
      ["Stan", "Oh my God, they killed Kenny!"],
      ["Kyle", "Dude, seriously?"],
      ["Kenny", "OMG!"],
      ["Butters", "Oh hamburgers…"],
      ["Randy", "I thought this was America!"],
    ];

    timer = setInterval(() => {
      const [a, b] = quotes[Math.floor(Math.random() * quotes.length)];
      $("author").value = a;
      $("body").value = b;
      send(a, b);
    }, 2000);
  };

  $("stopBtn").onclick = () => {
    clearInterval(timer);
    timer = null;
    $("autoBtn").disabled = false;
    $("stopBtn").disabled = true;
  };

  $("clearBtn").onclick = () => (logEl.innerHTML = "");

  // Initial health ping + keep-alive pings
  pingHealth();
  setInterval(pingHealth, 5000);
})();
