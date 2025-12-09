<!doctype html>
<html lang="lv">
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width,initial-scale=1" />
  <title>Gleznu galerija — Pievieno, skati, pārdod</title>

  <!-- Vienkāršs, mūsdienīgs stils -->
  <style>
    :root{
      --bg:#0f1724;
      --card:#0b1220;
      --muted:#98a2b3;
      --accent:#7c3aed;
      --glass: rgba(255,255,255,0.04);
      --white: #ffffff;
      --maxw:1100px;
    }
    *{box-sizing:border-box}
    body{
      margin:0;
      font-family: Inter, ui-sans-serif, system-ui, -apple-system, "Segoe UI", Roboto, "Helvetica Neue", Arial;
      background: linear-gradient(180deg,#071025 0%, #071827 60%);
      color:var(--white);
      -webkit-font-smoothing:antialiased;
      -moz-osx-font-smoothing:grayscale;
      padding:32px 16px;
    }

    .wrap{max-width:var(--maxw);margin:0 auto}

    header{
      display:flex;
      gap:16px;
      align-items:center;
      justify-content:space-between;
      margin-bottom:20px;
    }
    h1{font-size:20px;margin:0}
    p.lead{margin:0;color:var(--muted);font-size:13px}

    /* Controls area */
    .controls{
      display:grid;
      grid-template-columns: 1fr 360px;
      gap:20px;
      margin-bottom:22px;
    }

    /* Upload card */
    .card{
      background: linear-gradient(180deg, rgba(255,255,255,0.02), rgba(255,255,255,0.01));
      border: 1px solid rgba(255,255,255,0.04);
      padding:16px;
      border-radius:12px;
      box-shadow: 0 6px 30px rgba(2,6,23,0.6);
    }

    .uploader{
      display:flex;
      flex-direction:column;
      gap:10px;
    }

    .drop{
      border:2px dashed rgba(255,255,255,0.06);
      border-radius:10px;
      padding:16px;
      display:flex;
      gap:12px;
      align-items:center;
      justify-content:center;
      text-align:center;
      color:var(--muted);
      cursor:pointer;
      transition:background .15s, border-color .15s;
    }
    .drop.dragover{ background: rgba(255,255,255,0.02); border-color: rgba(124,58,237,0.7); color:var(--white) }

    .inputs{display:flex;flex-direction:column;gap:8px}
    input[type="text"], input[type="number"], textarea{
      width:100%;
      padding:10px 12px;
      border-radius:8px;
      border:1px solid rgba(255,255,255,0.04);
      background: rgba(255,255,255,0.02);
      color:var(--white);
      outline:none;
    }
    textarea{min-height:80px;resize:vertical}

    .row{display:flex;gap:8px}
    .btn{
      display:inline-flex;align-items:center;gap:8px;
      padding:10px 12px;border-radius:10px;border:none;cursor:pointer;
      background:linear-gradient(90deg,var(--accent),#4f46e5);color:var(--white);
      font-weight:600;
    }
    .btn.ghost{ background: transparent; border:1px solid rgba(255,255,255,0.04); color:var(--white) }
    .small{font-size:12px;color:var(--muted)}

    /* Gallery grid */
    .gallery{
      margin-top:8px;
      display:grid;
      grid-template-columns: repeat(auto-fill, minmax(240px, 1fr));
      gap:18px;
    }
    .art{
      background: linear-gradient(180deg, rgba(255,255,255,0.015), rgba(255,255,255,0.01));
      border-radius:12px;
      overflow:hidden;
      border:1px solid rgba(255,255,255,0.03);
      display:flex;
      flex-direction:column;
      transition:transform .15s;
    }
    .art:hover { transform: translateY(-6px) }

    .art-figure{ width:100%; height:220px; object-fit:cover; background:#0b1220; display:flex; align-items:center; justify-content:center; color:var(--muted) }
    .meta{padding:12px 14px; display:flex; flex-direction:column; gap:8px}
    .meta .title{font-weight:700;font-size:16px}
    .meta .desc{font-size:13px;color:var(--muted);min-height:36px}
    .price-row{display:flex;align-items:center;justify-content:space-between;gap:10px}
    .price{font-weight:800}
    .actions{display:flex;gap:8px;margin-top:8px}

    .action{ padding:8px 10px;border-radius:8px;border:none;background:transparent;color:var(--white); cursor:pointer;border:1px solid rgba(255,255,255,0.03) }
    .action.danger{border-color: rgba(255,30,60,0.2); color:#ffb3b3}

    /* Modal */
    .modal {
      position:fixed; inset:0; display:none; align-items:center; justify-content:center;
      background: rgba(2,6,23,0.7); z-index:60;
    }
    .modal.open{ display:flex }
    .modal-card{
      width:90%; max-width:760px; background:var(--card); border-radius:12px; padding:18px; border:1px solid rgba(255,255,255,0.04);
    }
    .modal-grid{ display:grid; grid-template-columns: 1fr 320px; gap:14px }
    .modal-img{ width:100%; height:420px; object-fit:contain; background:#061021; display:flex; align-items:center; justify-content:center}
    .modal .close{ float:right; background:transparent;border:none;color:var(--muted); cursor:pointer;font-size:18px }

    /* footer */
    footer{ margin-top:26px; text-align:center; color:var(--muted); font-size:13px }

    /* responsive */
    @media (max-width:920px){
      .controls{ grid-template-columns: 1fr; }
      .modal-grid{ grid-template-columns: 1fr; }
      .modal-img{ height:300px }
    }
  </style>
</head>
<body>
  <div class="wrap">
    <header>
      <div>
        <h1>Gleznu galerija</h1>
        <p class="lead">Pievieno gleznas, norādi cenas un pievieno PayPal linku, lai varētu pirkt.</p>
      </div>

      <div style="display:flex;gap:8px;align-items:center">
        <button class="btn" id="exportBtn" title="Eksportēt galeriju kā JSON">Eksportēt</button>
        <label class="btn ghost" style="padding:8px 10px;cursor:pointer">
          Importēt
          <input type="file" id="importFile" accept="application/json" style="display:none">
        </label>
      </div>
    </header>

    <div class="controls">
      <!-- Left: uploader -->
      <div class="card uploader">
        <div id="dropZone" class="drop">
          <div>
            <div style="font-weight:700">Pievieno attēlu</div>
            <div class="small">Velc šeit vai noklikšķini (PNG/JPG). Tiks saglabāts pārlūkā.</div>
          </div>
        </div>

        <div class="inputs">
          <input id="title" type="text" placeholder="Glezna — nosaukums" />
          <textarea id="desc" placeholder="Apraksts (piem., izmērs, tehnika)"></textarea>

          <div class="row">
            <input id="price" type="number" step="0.01" placeholder="Cena (€)" />
            <input id="currency" type="text" placeholder="Valūta (piem. EUR)" value="EUR" />
          </div>

          <input id="paypal" type="text" placeholder="PayPal vai cita maksājuma saite (piem. https://www.paypal.me/tevs/50)" />
          <div style="display:flex;gap:8px">
            <button class="btn" id="addBtn">Pievienot galerijai</button>
            <button class="btn ghost" id="clearBtn">Notīrīt laukus</button>
          </div>
          <div class="small" style="margin-top:6px">Dati tiek saglabāti tavā pārlūkā (localStorage). Eksportē, ja gribi backup.</div>
        </div>
      </div>

      <!-- Right: gallery preview -->
      <div>
        <div class="card">
          <div style="display:flex;align-items:center;justify-content:space-between">
            <div><strong>Galerija</strong></div>
            <div class="small" id="countLabel">0 gleznas</div>
          </div>

          <div id="gallery" class="gallery" style="margin-top:12px"></div>
        </div>
      </div>
    </div>

    <footer>Izveidots — vienkārša galerija ar lokālu saglabāšanu • Vari publicēt šo lapu ar GitHub Pages</footer>
  </div>

  <!-- Modal preview / edit -->
  <div id="modal" class="modal" aria-hidden="true">
    <div class="modal-card">
      <button class="close" id="closeModal">&times;</button>
      <div class="modal-grid">
        <div>
          <div id="modalImg" class="modal-img">Nav attēla</div>
        </div>
        <div>
          <div style="display:flex;flex-direction:column;gap:8px">
            <input id="m_title" type="text" placeholder="Nosaukums" />
            <textarea id="m_desc" placeholder="Apraksts"></textarea>
            <div class="row">
              <input id="m_price" type="number" step="0.01" placeholder="Cena" />
              <input id="m_currency" type="text" placeholder="Valūta" />
            </div>
            <input id="m_paypal" type="text" placeholder="PayPal vai maksājuma saite" />
            <div style="display:flex;gap:8px;margin-top:8px">
              <button class="btn" id="saveEdit">Saglabāt izmaiņas</button>
              <button class="btn ghost" id="cancelEdit">Atcelt</button>
            </div>
            <div class="small" style="margin-top:8px">Vari rediģēt vai pievienot maksājuma linku šeit.</div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- JavaScript: galerijas loģika -->
  <script>
    /* ========== Storage un helperi ========== */
    const STORAGE_KEY = 'gleznu_galerija_v1';

    function uid() {
      return Date.now().toString(36) + Math.random().toString(36).slice(2,8);
    }

    function saveAll(items){
      localStorage.setItem(STORAGE_KEY, JSON.stringify(items));
    }

    function loadAll(){
      try{
        const raw = localStorage.getItem(STORAGE_KEY);
        return raw ? JSON.parse(raw) : [];
      }catch(e){ return [] }
    }

    function formatPrice(price, currency){
      if(!price && price !== 0) return 'Cena nav norādīta';
      const num = Number(price);
      if(Number.isNaN(num)) return '—';
      const cur = currency ? currency.toUpperCase() : 'EUR';
      return num.toFixed(2) + ' ' + cur;
    }

    /* ========== DOM refs ========== */
    const dropZone = document.getElementById('dropZone');
    const addBtn = document.getElementById('addBtn');
    const clearBtn = document.getElementById('clearBtn');
    const galleryEl = document.getElementById('gallery');
    const countLabel = document.getElementById('countLabel');
    const exportBtn = document.getElementById('exportBtn');
    const importFile = document.getElementById('importFile');

    const inputs = {
      title: document.getElementById('title'),
      desc: document.getElementById('desc'),
      price: document.getElementById('price'),
      currency: document.getElementById('currency'),
      paypal: document.getElementById('paypal'),
    };

    let items = loadAll(); // items = [{id, title, desc, price, currency, paypal, imgDataUrl, createdAt}]

    /* ========== Drag & Drop + file open ========== */
    function prevent(e){ e.preventDefault(); e.stopPropagation(); }

    ['dragenter','dragover','dragleave','drop'].forEach(ev => {
      dropZone.addEventListener(ev, prevent);
    });

    dropZone.addEventListener('dragover', () => dropZone.classList.add('dragover'));
    dropZone.addEventListener('dragleave', () => dropZone.classList.remove('dragover'));
    dropZone.addEventListener('drop', (e) => {
      dropZone.classList.remove('dragover');
      const f = e.dataTransfer.files && e.dataTransfer.files[0];
      if(f) handleFile(f);
    });

    dropZone.addEventListener('click', () => {
      // create an invisible file input on the fly
      const fi = document.createElement('input'); fi.type='file'; fi.accept='image/*';
      fi.addEventListener('change', () => {
        if(fi.files && fi.files[0]) handleFile(fi.files[0]);
      });
      fi.click();
    });

    // compress/resize image to reasonable size (max 1600px)
    function handleFile(file){
      if(!file.type.startsWith('image/')) { alert('Izvēlies attēlu.'); return; }
      const reader = new FileReader();
      reader.onload = e => {
        resizeDataUrl(e.target.result, 1600, (resized) => {
          // fill image preview in uploader (optional)
          inputs._imgData = resized;
          dropZone.innerHTML = '<div style="font-weight:700">Attēls gatavs</div><div class="small">Nospied \"Pievienot galerijai\" vai pievieno citus laukus</div>';
        });
      };
      reader.readAsDataURL(file);
    }

    function resizeDataUrl(dataUrl, max, cb){
      const img = new Image();
      img.onload = () => {
        const ratio = Math.max(img.width / max, img.height / max, 1);
        const w = Math.round(img.width / ratio);
        const h = Math.round(img.height / ratio);
        const canvas = document.createElement('canvas');
        canvas.width = w; canvas.height = h;
        const ctx = canvas.getContext('2d');
        ctx.drawImage(img,0,0,w,h);
        const out = canvas.toDataURL('image/jpeg', 0.86);
        cb(out);
      };
      img.onerror = () => cb(dataUrl);
      img.src = dataUrl;
    }

    /* ========== Render galerija ========== */
    function render(){
      galleryEl.innerHTML = '';
      if(!items || items.length === 0){
        galleryEl.innerHTML = '<div class="small" style="padding:14px;color:var(--muted)">Galerija tukša — pievieno pirmo gleznu.</div>';
      } else {
        items.slice().reverse().forEach(item => { // newest first
          const el = document.createElement('div');
          el.className = 'art';
          el.innerHTML = `
            <div class="art-figure"><img src="${escapeHtml(item.imgDataUrl)}" alt="${escapeHtml(item.title)}" style="width:100%;height:100%;object-fit:cover"></div>
            <div class="meta">
              <div class="title">${escapeHtml(item.title || 'Bez nosaukuma')}</div>
              <div class="desc">${escapeHtml(item.desc || '')}</div>
              <div class="price-row">
                <div class="price">${formatPrice(item.price, item.currency)}</div>
              </div>
              <div class="actions">
                <button class="action" data-id="${item.id}" data-act="view">Skatīt</button>
                <button class="action" data-id="${item.id}" data-act="edit">Rediģēt</button>
                <button class="action" data-id="${item.id}" data-act="delete">Dzēst</button>
                <a class="action" style="margin-left:auto;background:linear-gradient(90deg,#06b6d4,#3b82f6);color:#031022;text-decoration:none" href="${escapeHtml(item.paypal || '#')}" target="_blank" rel="noopener" ${item.paypal ? '' : 'onclick="event.preventDefault()"'}>${item.paypal ? 'Pirkt' : 'Nav pirkšanas linka'}</a>
              </div>
            </div>
          `;
          galleryEl.appendChild(el);
        });
      }
      countLabel.textContent = items.length + (items.length===1 ? ' glezna' : ' gleznas');
    }

    /* ========== Actions: add, clear, edit, delete, view ========== */
    addBtn.addEventListener('click', () => {
      if(!inputs._imgData){ alert('Pievieno attēlu (velc vai noklikšķini uz zonas).'); return; }
      const title = inputs.title.value.trim();
      const desc = inputs.desc.value.trim();
      const price = inputs.price.value.trim();
      const currency = inputs.currency.value.trim() || 'EUR';
      const paypal = inputs.paypal.value.trim();

      const obj = {
        id: uid(),
        title, desc, price, currency, paypal,
        imgDataUrl: inputs._imgData,
        createdAt: new Date().toISOString()
      };
      items.push(obj);
      saveAll(items);
      // reset uploader UI
      inputs.title.value=''; inputs.desc.value=''; inputs.price.value=''; inputs.currency.value='EUR'; inputs.paypal.value=''; inputs._imgData = null;
      dropZone.innerHTML = `<div><div style="font-weight:700">Pievieno attēlu</div><div class="small">Velc šeit vai noklikšķini (PNG/JPG).</div></div>`;
      render();
    });

    clearBtn.addEventListener('click', () => {
      inputs.title.value=''; inputs.desc.value=''; inputs.price.value=''; inputs.currency.value='EUR'; inputs.paypal.value=''; inputs._imgData = null;
      dropZone.innerHTML = `<div><div style="font-weight:700">Pievieno attēlu</div><div class="small">Velc šeit vai noklikšķini (PNG/JPG).</div></div>`;
    });

    // delegated actions
    galleryEl.addEventListener('click', (e) => {
      const b = e.target.closest('button.action');
      if(!b) return;
      const id = b.getAttribute('data-id');
      const act = b.getAttribute('data-act');
      if(act === 'delete'){ if(confirm('Dzēst šo gleznu?')) { items = items.filter(it => it.id !== id); saveAll(items); render(); } }
      else if(act === 'edit'){ openModalForEdit(id); }
      else if(act === 'view'){ openModalForView(id); }
    });

    /* ========== Modal edit/view ========== */
    const modal = document.getElementById('modal');
    const closeModal = document.getElementById('closeModal');
    const cancelEdit = document.getElementById('cancelEdit');
    const saveEdit = document.getElementById('saveEdit');

    const m_title = document.getElementById('m_title');
    const m_desc = document.getElementById('m_desc');
    const m_price = document.getElementById('m_price');
    const m_currency = document.getElementById('m_currency');
    const m_paypal = document.getElementById('m_paypal');
    const modalImg = document.getElementById('modalImg');

    let editingId = null;

    function openModalForView(id){
      const it = items.find(x => x.id === id);
      if(!it) return;
      editingId = null;
      modal.classList.add('open');
      modal.setAttribute('aria-hidden','false');
      modalImg.innerHTML = `<img src="${escapeHtml(it.imgDataUrl)}" style="width:100%;height:100%;object-fit:contain">`;
      m_title.value = it.title || '';
      m_desc.value = it.desc || '';
      m_price.value = it.price || '';
      m_currency.value = it.currency || 'EUR';
      m_paypal.value = it.paypal || '';
      // disable inputs for pure view
      [m_title,m_desc,m_price,m_currency,m_paypal].forEach(i => i.disabled = true);
      saveEdit.style.display = 'none';
    }

    function openModalForEdit(id){
      const it = items.find(x => x.id === id);
      if(!it) return;
      editingId = id;
      modal.classList.add('open');
      modal.setAttribute('aria-hidden','false');
      modalImg.innerHTML = `<img src="${escapeHtml(it.imgDataUrl)}" style="width:100%;height:100%;object-fit:contain">`;
      m_title.value = it.title || '';
      m_desc.value = it.desc || '';
      m_price.value = it.price || '';
      m_currency.value = it.currency || 'EUR';
      m_paypal.value = it.paypal || '';
      [m_title,m_desc,m_price,m_currency,m_paypal].forEach(i => i.disabled = false);
      saveEdit.style.display = '';
    }

    function closeModalFn(){
      modal.classList.remove('open');
      modal.setAttribute('aria-hidden','true');
      editingId = null;
    }

    closeModal.addEventListener('click', closeModalFn);
    cancelEdit.addEventListener('click', closeModalFn);
    modal.addEventListener('click', (e) => { if(e.target === modal) closeModalFn(); });

    saveEdit.addEventListener('click', () => {
      if(!editingId) return;
      const it = items.find(x => x.id === editingId);
      if(!it) return;
      it.title = m_title.value.trim();
      it.desc = m_desc.value.trim();
      it.price = m_price.value.trim();
      it.currency = m_currency.value.trim() || 'EUR';
      it.paypal = m_paypal.value.trim();
      saveAll(items);
      render();
      closeModalFn();
    });

    /* ========== Export / Import JSON ========== */
    exportBtn.addEventListener('click', () => {
      const data = JSON.stringify(items, null, 2);
      const blob = new Blob([data], {type:'application/json'});
      const url = URL.createObjectURL(blob);
      const a = document.createElement('a');
      a.href = url; a.download = 'gleznu-galerija.json';
      document.body.appendChild(a); a.click(); a.remove();
      URL.revokeObjectURL(url);
    });

    importFile.addEventListener('change', (e) => {
      const f = e.target.files[0];
      if(!f) return;
      const r = new FileReader();
      r.onload = ev => {
        try{
          const parsed = JSON.parse(ev.target.result);
          if(Array.isArray(parsed)){
            // basic validation - must have id and imgDataUrl
            const ok = parsed.every(it => it.id && it.imgDataUrl);
            if(!ok) throw new Error('Nederīgs JSON formāts');
            items = parsed;
            saveAll(items);
            render();
            alert('Importēts veiksmīgi');
          } else {
            throw new Error('JSON nav masīvs');
          }
        }catch(err){
          alert('Importēšanas kļūda: ' + (err.message || err));
        }
      };
      r.readAsText(f);
      e.target.value = '';
    });

    /* ========== Utilities ========== */
    function escapeHtml(s){
      if(s === null || s === undefined) return '';
      return String(s).replaceAll('&','&amp;').replaceAll('<','&lt;').replaceAll('>','&gt;').replaceAll('"','&quot;');
    }

    /* ========== Init ========== */
    render();

  </script>
</body>
</html>
