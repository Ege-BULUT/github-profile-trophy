package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const landingPageHTML = `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta name="theme-color" content="#07111d">
  <meta name="description" content="Generate a live GitHub Profile Trophy card with a polished preview.">
  <link rel="canonical" href="https://github-torphy.vercel.app/">
  <meta property="og:type" content="website">
  <meta property="og:site_name" content="GitHub Trophy Lab">
  <meta property="og:url" content="https://github-torphy.vercel.app/">
  <meta property="og:title" content="GitHub Trophy Lab">
  <meta property="og:description" content="Generate a live GitHub Profile Trophy card with a polished preview.">
  <meta name="twitter:card" content="summary">
  <meta name="google-site-verification" content="s66rd8O_x1saL6PLGpSpIPeHmGazHyTo27NTPkK-GJg">
  <title>GitHub Trophy Lab</title>
  <style>
    :root {
      color-scheme: dark;
      --ink: #f4f8fb;
      --muted: #9baebe;
      --quiet: #718294;
      --line: rgba(214, 237, 246, .14);
      --glass: rgba(17, 31, 47, .62);
      --glass-bright: rgba(40, 63, 80, .58);
      --mint: #8df0cb;
      --mint-strong: #4ad6ab;
      --danger: #ff9caa;
      --radius-xl: 28px;
      --radius-lg: 20px;
      --shadow: 0 26px 80px rgba(0, 0, 0, .34);
    }

    * { box-sizing: border-box; }
    html { background: #07111d; }
    body {
      min-width: 320px;
      min-height: 100vh;
      margin: 0;
      overflow-x: clip;
      color: var(--ink);
      font-family: Inter, ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif;
      background:
        radial-gradient(circle at 10% 6%, rgba(42, 151, 191, .30), transparent 30rem),
        radial-gradient(circle at 92% 14%, rgba(77, 207, 168, .22), transparent 26rem),
        radial-gradient(circle at 72% 88%, rgba(64, 120, 255, .14), transparent 30rem),
        linear-gradient(145deg, #07111d 0%, #0b1826 52%, #07131c 100%);
    }

    body::before {
      position: fixed;
      inset: 0;
      z-index: -1;
      pointer-events: none;
      opacity: .35;
      content: "";
      background-image: linear-gradient(rgba(190, 228, 236, .04) 1px, transparent 1px), linear-gradient(90deg, rgba(190, 228, 236, .04) 1px, transparent 1px);
      background-size: 48px 48px;
      mask-image: linear-gradient(to bottom, black, transparent 80%);
    }

    a { color: inherit; }
    button, input, select { font: inherit; }
    button, a { -webkit-tap-highlight-color: transparent; }
    button { cursor: pointer; }
    :focus-visible { outline: 3px solid rgba(141, 240, 203, .8); outline-offset: 3px; }

    .shell { width: min(1180px, calc(100% - 40px)); margin: 0 auto; }
    .topbar { display: flex; flex-wrap: wrap; align-items: center; justify-content: space-between; gap: 16px; padding: 28px 0 62px; }
    .brand { display: inline-flex; align-items: center; gap: 12px; text-decoration: none; }
    .brand-mark { display: grid; width: 38px; height: 38px; place-items: center; border: 1px solid rgba(141, 240, 203, .42); border-radius: 12px; color: var(--mint); background: rgba(141, 240, 203, .1); box-shadow: inset 0 1px rgba(255,255,255,.18); }
    .brand-name { font-size: 14px; font-weight: 700; letter-spacing: .02em; }
    .brand-sub { display: block; margin-top: 3px; color: var(--quiet); font-size: 11px; letter-spacing: .11em; text-transform: uppercase; }
    .status-chip { display: inline-flex; align-items: center; gap: 8px; min-height: 34px; padding: 0 12px; border: 1px solid var(--line); border-radius: 999px; color: var(--muted); background: rgba(255,255,255,.04); font-size: 12px; }
    .status-dot { width: 7px; height: 7px; border-radius: 50%; background: var(--mint-strong); box-shadow: 0 0 14px var(--mint-strong); }

    .hero { display: grid; grid-template-columns: minmax(0, 1.03fr) minmax(0, .97fr); gap: 36px; align-items: end; padding-bottom: 30px; }
    .eyebrow { display: flex; align-items: center; gap: 10px; margin: 0 0 18px; color: var(--mint); font-size: 12px; font-weight: 700; letter-spacing: .16em; text-transform: uppercase; }
    .eyebrow::before { width: 28px; height: 1px; content: ""; background: var(--mint); }
    h1 { max-width: 680px; margin: 0; font-size: clamp(42px, 6vw, 78px); line-height: .96; letter-spacing: -.065em; }
    .hero-copy { max-width: 540px; margin: 24px 0 0; color: var(--muted); font-size: 17px; line-height: 1.65; }
    .hero-note { align-self: end; max-width: 320px; margin: 0 0 6px auto; padding: 14px 16px; border: 1px solid var(--line); border-radius: 16px; color: var(--muted); background: rgba(255,255,255,.05); backdrop-filter: blur(14px); -webkit-backdrop-filter: blur(14px); box-shadow: inset 0 1px rgba(255,255,255,.10); font-size: 13px; line-height: 1.55; }

    .workspace { display: grid; grid-template-columns: minmax(0, .92fr) minmax(0, 1.08fr); gap: 16px; padding-bottom: 24px; }
    .glass { border: 1px solid rgba(255,255,255,.16); border-radius: var(--radius-xl); background: linear-gradient(165deg, rgba(255,255,255,.10), rgba(255,255,255,.04) 34%, rgba(10,22,35,.50)); box-shadow: var(--shadow), inset 0 1px rgba(255,255,255,.14), inset 0 -1px rgba(255,255,255,.05); backdrop-filter: blur(26px) saturate(150%); -webkit-backdrop-filter: blur(26px) saturate(150%); }
    @supports not (backdrop-filter: blur(1px)) { .glass { background: #142536; } }
    .panel { min-width: 0; padding: 26px; }
    .panel-heading { display: flex; flex-wrap: wrap; justify-content: space-between; gap: 10px; margin-bottom: 24px; }
    .panel-title { margin: 0; font-size: 16px; letter-spacing: -.02em; }
    .panel-kicker { color: var(--quiet); font-size: 12px; }
    .form-stack { display: grid; gap: 18px; }
    label { display: block; margin-bottom: 8px; color: var(--muted); font-size: 12px; font-weight: 650; }
    input, select { width: 100%; min-height: 48px; padding: 0 14px; border: 1px solid rgba(214, 237, 246, .16); border-radius: 14px; color: var(--ink); background: rgba(3, 13, 24, .48); transition: border-color .2s ease, background .2s ease, box-shadow .2s ease; }
    input::placeholder { color: #637688; }
    input:hover, select:hover { border-color: rgba(214, 237, 246, .28); }
    input:focus, select:focus { border-color: rgba(141, 240, 203, .72); background: rgba(3, 13, 24, .68); box-shadow: 0 0 0 4px rgba(141, 240, 203, .1); outline: none; }
    select { appearance: none; background-image: linear-gradient(45deg, transparent 50%, var(--muted) 50%), linear-gradient(135deg, var(--muted) 50%, transparent 50%); background-position: calc(100% - 18px) 21px, calc(100% - 13px) 21px; background-repeat: no-repeat; background-size: 5px 5px, 5px 5px; }
    option { color: #0b1826; }
    .field-grid { display: grid; grid-template-columns: 1fr 1fr; gap: 12px; }
    .toggle-row { display: flex; flex-wrap: wrap; gap: 10px; }
    .toggle { display: inline-flex; align-items: center; gap: 9px; min-height: 42px; padding: 0 12px; border: 1px solid rgba(214, 237, 246, .14); border-radius: 12px; color: var(--muted); background: rgba(255,255,255,.035); font-size: 12px; }
    .toggle input { width: 16px; min-height: 16px; accent-color: var(--mint-strong); }
    .primary { display: inline-flex; align-items: center; justify-content: center; gap: 10px; width: 100%; min-height: 50px; margin-top: 2px; border: 0; border-radius: 14px; color: #052019; background: linear-gradient(135deg, var(--mint), #b0ffe4); box-shadow: 0 12px 30px rgba(74, 214, 171, .2); font-weight: 800; transition: transform .2s ease, box-shadow .2s ease; }
    .primary:hover { transform: translateY(-2px); box-shadow: 0 16px 36px rgba(74, 214, 171, .32); }
    .primary:active { transform: translateY(0); }
    .primary:disabled { opacity: .45; cursor: not-allowed; transform: none; box-shadow: none; }

    .preview-panel { display: flex; min-height: 100%; flex-direction: column; padding: 0; overflow: hidden; }
    .preview-head { display: flex; flex-wrap: wrap; align-items: center; justify-content: space-between; gap: 12px; padding: 26px 26px 0; }
    .preview-state { color: var(--quiet); font-size: 12px; white-space: nowrap; overflow: hidden; text-overflow: ellipsis; }
    .preview-state[data-tone="error"] { color: var(--danger); }
    .preview-stage { display: grid; min-height: 332px; flex: 1; margin: 22px; place-items: center; padding: 28px; border: 1px solid rgba(214, 237, 246, .1); border-radius: var(--radius-lg); background: radial-gradient(circle at 50% 35%, rgba(85, 176, 190, .12), transparent 52%), rgba(3, 13, 24, .48); }
    .preview-stage img { display: block; max-width: 100%; height: auto; opacity: 0; transition: opacity .3s ease; }
    .preview-stage img.is-visible { opacity: 1; }
    .preview-empty { max-width: 270px; color: var(--quiet); text-align: center; line-height: 1.6; }
    .preview-empty strong { display: block; margin-bottom: 8px; color: var(--muted); font-size: 14px; }
    .preview-empty svg { margin-bottom: 16px; color: var(--mint); }
    .preview-footer { display: flex; flex-wrap: wrap; align-items: center; justify-content: space-between; gap: 12px; padding: 0 26px 26px; }
    .preview-footer small { color: var(--quiet); font-size: 11px; }
    .copy-button { min-height: 40px; padding: 0 13px; border: 1px solid rgba(214, 237, 246, .16); border-radius: 11px; color: var(--muted); background: rgba(255,255,255,.04); font-size: 12px; transition: background .2s ease, color .2s ease; }
    .copy-button:hover { color: var(--ink); background: rgba(255,255,255,.1); }
    .copy-button:disabled { cursor: not-allowed; opacity: .45; }
    .copy-button.toast { color: var(--mint-strong); border-color: rgba(141, 240, 203, .3); }
    .advanced-toggle { display: inline-flex; align-items: center; gap: 8px; width: 100%; padding: 10px 0; border: 0; background: none; color: var(--quiet); font-size: 12px; font-weight: 650; cursor: pointer; transition: color .2s ease; }
    .advanced-toggle:hover { color: var(--muted); }
    .advanced-toggle svg { transition: transform .2s ease; }
    .advanced-toggle.open svg { transform: rotate(180deg); }
    .advanced-section { display: none; padding-top: 14px; margin-top: 14px; border-top: 1px solid var(--line); }
    .advanced-section.open { display: block; }
    .advanced-section .field-grid { margin-top: 0; }
    .footer { display: flex; flex-wrap: wrap; align-items: center; justify-content: space-between; gap: 14px; padding: 12px 0 30px; color: var(--quiet); font-size: 12px; }
    .footer a { color: var(--muted); text-underline-offset: 3px; }
    .footer a:hover { color: var(--mint); }
    .sr-only { position: absolute; width: 1px; height: 1px; padding: 0; margin: -1px; overflow: hidden; clip: rect(0, 0, 0, 0); white-space: nowrap; border: 0; }

    @media (max-width: 840px) {
      .topbar { padding-bottom: 26px; }
      .hero, .workspace { grid-template-columns: 1fr; }
      .hero { padding-bottom: 18px; }
      .hero-copy { margin-top: 16px; }
      .hero-note { margin: 14px 0 0; max-width: none; }
      .preview-stage { min-height: 280px; }
    }
    @media (max-width: 520px) {
      .shell { width: min(100% - 28px, 1180px); }
      .topbar { padding-top: 18px; padding-bottom: 22px; }
      .form-stack { gap: 14px; }
      .status-chip { width: 100%; justify-content: center; }
      h1 { font-size: clamp(40px, 13vw, 58px); }
      .hero-copy { font-size: 15px; }
      .panel { padding: 20px; }
      .field-grid { grid-template-columns: 1fr; }
      .preview-head { padding: 20px 20px 0; }
      .preview-stage { min-height: 250px; margin: 18px; padding: 18px; }
      .preview-footer { padding: 0 20px 20px; }
    }
    @media (prefers-reduced-motion: reduce) {
      *, *::before, *::after { scroll-behavior: auto !important; transition-duration: .01ms !important; animation-duration: .01ms !important; animation-iteration-count: 1 !important; }
    }
  </style>
</head>
<body>
  <main class="shell">
    <header class="topbar">
      <a class="brand" href="/" aria-label="GitHub Trophy Lab home">
        <span class="brand-mark" aria-hidden="true"><svg width="21" height="21" viewBox="0 0 24 24" fill="none"><path d="M7 4h10v4.5a5 5 0 0 1-10 0V4Z" stroke="currentColor" stroke-width="1.7"/><path d="M7 6H4.8A1.8 1.8 0 0 0 3 7.8v.4a3.4 3.4 0 0 0 3.4 3.4M17 6h2.2A1.8 1.8 0 0 1 21 7.8v.4a3.4 3.4 0 0 1-3.4 3.4M12 13.5V18M8.5 21h7M9 18h6" stroke="currentColor" stroke-linecap="round" stroke-width="1.7"/></svg></span>
        <span><span class="brand-name">GitHub Trophy Lab</span><span class="brand-sub">Profile card studio</span></span>
      </a>
      <span class="status-chip"><span class="status-dot" aria-hidden="true"></span> SVG renderer online</span>
    </header>

    <section class="hero" aria-labelledby="page-title">
      <div>
        <p class="eyebrow">Build your signal</p>
        <h1 id="page-title">Your work,<br><em>worth displaying.</em></h1>
        <p class="hero-copy">Turn GitHub activity into a trophy card that feels at home in your profile README. Tune the layout, see the result, copy the embed.</p>
      </div>
      <p class="hero-note">A self-hosted, upstream-compatible renderer for people who care about the details.</p>
    </section>

    <section class="workspace" aria-label="Trophy card builder">
      <form class="glass panel" id="trophy-form" action="/" method="get">
        <div class="panel-heading"><div><h2 class="panel-title">Card settings</h2><span class="panel-kicker">Changes preview automatically</span></div><span class="panel-kicker" id="form-hint">Required: username</span></div>
        <div class="form-stack">
          <div>
            <label for="username">GitHub username</label>
            <input id="username" name="username" type="text" autocomplete="off" placeholder="e.g. ege-bulut" required aria-describedby="username-help">
            <span class="sr-only" id="username-help">Enter a public GitHub username to generate a trophy card.</span>
          </div>
          <div>
            <label for="theme">Visual theme</label>
            <select id="theme" name="theme">
              <option value="default">Default</option><option value="onedark">One Dark</option><option value="dracula">Dracula</option><option value="nord">Nord</option><option value="gruvbox">Gruvbox</option><option value="radical">Radical</option><option value="tokyonight">Tokyo Night</option><option value="aura">Aura</option>
            </select>
          </div>
          <div class="field-grid">
            <div><label for="column">Columns</label><select id="column" name="column"><option value="-1">Adaptive</option><option value="4">4</option><option value="5">5</option><option value="6">6</option><option value="7">7</option><option value="8">8</option></select></div>
            <div><label for="margin-w">Card gap</label><select id="margin-w" name="margin-w"><option value="25">25 px</option><option value="15">15 px</option><option value="10">10 px</option><option value="0">None</option></select></div>
          </div>
          <div class="toggle-row">
             <label class="toggle"><input id="no-bg" name="no-bg" type="checkbox" value="true" checked> Transparent background</label>
             <label class="toggle"><input id="no-frame" name="no-frame" type="checkbox" value="true"> Hide frames</label>
           </div>
           <button class="advanced-toggle" id="advanced-toggle" type="button" aria-expanded="false" aria-controls="advanced-section">
             <span>Advanced settings</span>
             <svg width="12" height="12" viewBox="0 0 12 12" fill="none" aria-hidden="true"><path d="M3 5h6L6 9z" fill="currentColor"/></svg>
           </button>
           <div class="advanced-section" id="advanced-section">
             <div class="field-grid">
               <div><label for="rank">Rank filter</label><input id="rank" name="rank" type="text" placeholder="e.g. S,SS,SSS or -A,-B" autocomplete="off"></div>
               <div><label for="row">Max rows</label><input id="row" name="row" type="number" min="1" max="20" placeholder="e.g. 5" autocomplete="off"></div>
             </div>
             <div class="field-grid" style="margin-top: 12px;">
               <div><label for="margin-h">Row gap</label><input id="margin-h" name="margin-h" type="number" min="0" max="100" placeholder="e.g. 25" autocomplete="off"></div>
               <div><label for="title">Title filter</label><input id="title" name="title" type="text" placeholder="e.g. Followers or -Stars,-Forks" autocomplete="off"></div>
             </div>
           </div>
           <button class="primary" type="submit" disabled><span>Generate trophy card</span><span aria-hidden="true">↗</span></button>
        </div>
      </form>

      <section class="glass preview-panel" aria-labelledby="preview-title">
        <div class="preview-head"><div><h2 class="panel-title" id="preview-title">Live preview</h2><span class="panel-kicker">Rendered from your endpoint</span></div><span class="preview-state" id="preview-state" role="status" aria-live="polite">Waiting for a username</span></div>
        <div class="preview-stage" id="preview-stage">
          <div class="preview-empty" id="preview-empty"><svg width="42" height="42" viewBox="0 0 42 42" fill="none" aria-hidden="true"><circle cx="21" cy="21" r="19" stroke="currentColor" stroke-width="1.5"/><path d="M13 23.5 18.2 28 29 15" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.8"/></svg><strong>Your trophy card will appear here</strong>Type a username and the renderer will build a fresh SVG preview.</div>
          <img id="preview-image" alt="GitHub trophy card preview" decoding="async">
        </div>
        <div class="preview-footer"><small id="preview-url">No preview generated yet</small><button class="copy-button" id="copy-button" type="button" disabled>Copy Markdown</button></div>
      </section>
    </section>

     <footer class="footer"><span>Designed for profile README embeds.</span><span><a href="/docs" target="_self">Docs</a><span aria-hidden="true"> · </span><a href="/reference" target="_self">Reference</a><span aria-hidden="true"> · </span><a href="https://github.com/soulteary/github-profile-trophy" target="_blank" rel="noreferrer">View upstream</a><span aria-hidden="true"> · </span><a href="https://github.com/ege-bulut" target="_blank" rel="noreferrer">ege-bulut</a></span></footer>
  </main>
  <script>
    (() => {
      const form = document.querySelector('#trophy-form');
      const username = document.querySelector('#username');
      const image = document.querySelector('#preview-image');
      const empty = document.querySelector('#preview-empty');
      const state = document.querySelector('#preview-state');
      const previewUrl = document.querySelector('#preview-url');
      const copyButton = document.querySelector('#copy-button');
      const submitButton = form.querySelector('.primary');
      let timer;
      let activeUrl = '';

      const setState = (message, tone = '') => {
        state.textContent = message;
        state.dataset.tone = tone;
      };

      const buildUrl = () => {
        const data = new FormData(form);
        const params = new URLSearchParams();
        for (const [key, value] of data.entries()) {
          if (value) params.set(key, String(value));
        }
        if (!data.has('no-bg')) params.delete('no-bg');
        if (!data.has('no-frame')) params.delete('no-frame');
        return '/?' + params.toString();
      };

      const renderPreview = async () => {
        const name = username.value.trim();
        if (name.length < 3) {
          setState('Waiting for a username');
          return;
        }
        const url = buildUrl();
        setState('Rendering preview…');
        image.classList.remove('is-visible');
        try {
          const response = await fetch(url, { headers: { Accept: 'image/svg+xml' } });
          if (!response.ok) throw new Error(response.status === 404 ? 'GitHub user not found' : response.status === 500 ? 'Renderer needs a GITHUB_TOKEN1 environment variable' : 'Renderer could not create this card');
          const svg = await response.text();
          image.src = 'data:image/svg+xml;charset=utf-8,' + encodeURIComponent(svg);
          image.onload = () => { empty.hidden = true; image.classList.add('is-visible'); };
          activeUrl = new URL(url, window.location.href).href;
          previewUrl.textContent = activeUrl;
          copyButton.disabled = false;
          setState('Preview ready');
        } catch (error) {
          empty.hidden = false;
          setState(error.message || 'Preview failed', 'error');
          previewUrl.textContent = 'Check the username or renderer token';
          copyButton.disabled = true;
        }
      };

      const schedulePreview = () => {
        clearTimeout(timer);
        timer = setTimeout(renderPreview, 700);
      };

       const syncSubmit = () => { submitButton.disabled = username.value.trim().length < 3; };

       const advancedToggle = document.querySelector('#advanced-toggle');
       const advancedSection = document.querySelector('#advanced-section');
       advancedToggle.addEventListener('click', () => {
         const isOpen = advancedSection.classList.toggle('open');
         advancedToggle.classList.toggle('open', isOpen);
         advancedToggle.setAttribute('aria-expanded', String(isOpen));
       });

       form.addEventListener('submit', (event) => { event.preventDefault(); clearTimeout(timer); renderPreview(); });
       form.addEventListener('input', () => { syncSubmit(); schedulePreview(); });
       form.addEventListener('change', () => { syncSubmit(); schedulePreview(); });
       syncSubmit();
       copyButton.addEventListener('click', async () => {
         if (!activeUrl) return;
         const markdown = '[![trophy](' + activeUrl + ')](https://github.com/' + encodeURIComponent(username.value.trim()) + ')';
         await navigator.clipboard.writeText(markdown);
         copyButton.textContent = 'Copied';
         copyButton.classList.add('toast');
         setTimeout(() => { copyButton.textContent = 'Copy Markdown'; copyButton.classList.remove('toast'); }, 1600);
       });
    })();
  </script>
</body>
</html>`

const docsPageHTML = `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta name="theme-color" content="#07111d">
  <title>GitHub Trophy Lab — Docs</title>
  <meta name="description" content="How GitHub Trophy Lab works: the trophy endpoint, its query parameters and running your own renderer.">
  <link rel="canonical" href="https://github-torphy.vercel.app/docs">
  <style>
    :root { color-scheme: dark; --ink: #f4f8fb; --muted: #9baebe; --quiet: #718294; --line: rgba(214, 237, 246, .14); --mint: #8df0cb; --radius-xl: 28px; --shadow: 0 26px 80px rgba(0,0,0,.34); }
    * { box-sizing: border-box; }
    html { background: #07111d; }
    body { min-width: 320px; min-height: 100vh; margin: 0; color: var(--ink); font-family: Inter, ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif; background: radial-gradient(circle at 10% 6%, rgba(42,151,191,.30), transparent 30rem), radial-gradient(circle at 92% 14%, rgba(77,207,168,.22), transparent 26rem), linear-gradient(145deg, #07111d 0%, #0b1826 52%, #07131c 100%); }
    .shell { width: min(1180px, calc(100% - 40px)); margin: 0 auto; padding: 60px 0 80px; }
    .topbar { display: flex; align-items: center; justify-content: space-between; padding-bottom: 40px; }
    .brand { display: inline-flex; align-items: center; gap: 12px; text-decoration: none; color: inherit; }
    .brand-mark { display: grid; width: 38px; height: 38px; place-items: center; border: 1px solid rgba(141,240,203,.42); border-radius: 12px; color: var(--mint); background: rgba(141,240,203,.1); }
    .brand-name { font-size: 14px; font-weight: 700; }
    .brand-sub { display: block; margin-top: 3px; color: var(--quiet); font-size: 11px; letter-spacing: .11em; text-transform: uppercase; }
    h1 { font-size: clamp(36px, 5vw, 64px); line-height: .96; letter-spacing: -.065em; margin: 0 0 20px; }
    .eyebrow { display: flex; align-items: center; gap: 10px; margin: 0 0 18px; color: var(--mint); font-size: 12px; font-weight: 700; letter-spacing: .16em; text-transform: uppercase; }
    .eyebrow::before { width: 28px; height: 1px; content: ""; background: var(--mint); }
    .content { line-height: 1.75; color: var(--muted); font-size: 16px; }
    .content h2 { color: var(--ink); font-size: 1.4em; margin: 2em 0 .5em; }
    .content h3 { color: var(--ink); font-size: 1.15em; margin: 1.5em 0 .4em; }
    .content code { background: rgba(255,255,255,.06); padding: 2px 7px; border-radius: 6px; font-size: .9em; color: var(--mint); }
    .content pre { background: rgba(3,13,24,.6); padding: 16px 20px; border-radius: 14px; overflow-x: auto; border: 1px solid var(--line); }
    .content pre code { background: none; padding: 0; color: var(--ink); }
    .content a { color: var(--mint); text-underline-offset: 3px; }
    .content ul { padding-left: 1.2em; }
    .content li { margin-bottom: .4em; }
    .badge { display: inline-block; padding: 2px 9px; border-radius: 999px; background: rgba(141,240,203,.1); color: var(--mint); font-size: 12px; font-weight: 650; }
    .footer { padding-top: 40px; border-top: 1px solid var(--line); color: var(--quiet); font-size: 12px; display: flex; justify-content: space-between; flex-wrap: wrap; gap: 12px; }
    .footer a { color: var(--muted); }
    .footer a:hover { color: var(--mint); }
    @media (prefers-reduced-motion: reduce) { *, *::before, *::after { scroll-behavior: auto !important; transition-duration: .01ms !important; animation-duration: .01ms !important; } }
  </style>
</head>
<body>
  <main class="shell">
    <header class="topbar">
      <a class="brand" href="/" aria-label="GitHub Trophy Lab home">
        <span class="brand-mark"><svg width="21" height="21" viewBox="0 0 24 24" fill="none"><path d="M7 4h10v4.5a5 5 0 0 1-10 0V4Z" stroke="currentColor" stroke-width="1.7"/><path d="M7 6H4.8A1.8 1.8 0 0 0 3 7.8v.4a3.4 3.4 0 0 0 3.4 3.4M17 6h2.2A1.8 1.8 0 0 1 21 7.8v.4a3.4 3.4 0 0 1-3.4 3.4M12 13.5V18M8.5 21h7M9 18h6" stroke="currentColor" stroke-linecap="round" stroke-width="1.7"/></svg></span>
        <span><span class="brand-name">GitHub Trophy Lab</span><span class="brand-sub">Profile card studio</span></span>
      </a>
      <span class="badge">Docs</span>
    </header>
    <div class="content">
      <p class="eyebrow">Documentation</p>
      <h1>How it works</h1>
      <p>This service renders GitHub profile trophy cards as SVG — the same engine that powers <a href="https://github.com/soulteary/github-profile-trophy" target="_blank" rel="noreferrer">soulteary/github-profile-trophy</a>, self-hosted on Vercel.</p>
      <h2>Quick start</h2>
      <p>Paste the following into your GitHub README to embed a trophy card:</p>
      <pre><code>[![trophy](https://github-torphy.vercel.app/?username=YOUR_USERNAME)]
(https://github.com/YOUR_USERNAME)</code></pre>
      <h2>Query parameters</h2>
      <table style="width:100%;border-collapse:collapse;margin:1em 0;">
        <thead><tr style="border-bottom:1px solid var(--line);text-align:left"><th style="padding:8px 12px;color:var(--ink)">Parameter</th><th style="padding:8px 12px;color:var(--ink)">Values</th><th style="padding:8px 12px;color:var(--ink)">Default</th><th style="padding:8px 12px;color:var(--ink)">Description</th></tr></thead>
        <tbody>
          <tr style="border-bottom:1px solid var(--line)"><td style="padding:8px 12px"><code>username</code></td><td style="padding:8px 12px">string</td><td style="padding:8px 12px">—</td><td style="padding:8px 12px">GitHub username (required)</td></tr>
          <tr style="border-bottom:1px solid var(--line)"><td style="padding:8px 12px"><code>theme</code></td><td style="padding:8px 12px">string</td><td style="padding:8px 12px">default</td><td style="padding:8px 12px">Visual theme (default, onedark, dracula, nord, gruvbox, radical, tokyonight, aura)</td></tr>
          <tr style="border-bottom:1px solid var(--line)"><td style="padding:8px 12px"><code>column</code></td><td style="padding:8px 12px">integer</td><td style="padding:8px 12px">-1</td><td style="padding:8px 12px">Max columns; -1 = adaptive</td></tr>
          <tr style="border-bottom:1px solid var(--line)"><td style="padding:8px 12px"><code>row</code></td><td style="padding:8px 12px">integer</td><td style="padding:8px 12px">-1</td><td style="padding:8px 12px">Max rows; -1 = adaptive</td></tr>
          <tr style="border-bottom:1px solid var(--line)"><td style="padding:8px 12px"><code>margin-w</code></td><td style="padding:8px 12px">integer</td><td style="padding:8px 12px">25</td><td style="padding:8px 12px">Horizontal margin between trophies (px)</td></tr>
          <tr style="border-bottom:1px solid var(--line)"><td style="padding:8px 12px"><code>margin-h</code></td><td style="padding:8px 12px">integer</td><td style="padding:8px 12px">25</td><td style="padding:8px 12px">Vertical margin between trophies (px)</td></tr>
          <tr style="border-bottom:1px solid var(--line)"><td style="padding:8px 12px"><code>no-bg</code></td><td style="padding:8px 12px">boolean</td><td style="padding:8px 12px">true</td><td style="padding:8px 12px">Transparent background</td></tr>
          <tr style="border-bottom:1px solid var(--line)"><td style="padding:8px 12px"><code>no-frame</code></td><td style="padding:8px 12px">boolean</td><td style="padding:8px 12px">false</td><td style="padding:8px 12px">Hide trophy frames</td></tr>
          <tr style="border-bottom:1px solid var(--line)"><td style="padding:8px 12px"><code>title</code></td><td style="padding:8px 12px">string</td><td style="padding:8px 12px">—</td><td style="padding:8px 12px">Filter by trophy titles (comma-separated; prefix with <code>-</code> to exclude)</td></tr>
          <tr><td style="padding:8px 12px"><code>rank</code></td><td style="padding:8px 12px">string</td><td style="padding:8px 12px">—</td><td style="padding:8px 12px">Filter by ranks (comma-separated; prefix with <code>-</code> to exclude). Values: SECRET, SSS, SS, S, AAA, AA, A, B, C</td></tr>
        </tbody>
      </table>
      <h2>Examples</h2>
      <pre><code># Minimal trophy card
https://github-torphy.vercel.app/?username=ege-bulut

# With custom layout and transparent background
https://github-torphy.vercel.app/?username=ege-bulut&theme=dracula&column=7&margin-w=15&no-bg=true

# Filter by title and rank
https://github-torphy.vercel.app/?username=ege-bulut&title=Followers,Stars&rank=S,SSS

# Exclude certain titles
https://github-torphy.vercel.app/?username=ege-bulut&title=-Stars,-Forks</code></pre>
      <h2>Authentication</h2>
      <p>The renderer requires a <code>GITHUB_TOKEN1</code> environment variable (GitHub Personal Access Token) to fetch profile and trophy data via the GitHub API.</p>
      <h2>Self-hosting</h2>
      <p>This is a fork of <a href="https://github.com/soulteary/github-profile-trophy" target="_blank" rel="noreferrer">soulteary/github-profile-trophy</a> with a polished studio UI and Vercel deployment.</p>
      <ul>
        <li>Clone the repo and set <code>GITHUB_TOKEN1</code></li>
        <li>Deploy to Vercel or run locally with <code>go run ./cmd/server</code></li>
      </ul>
    </div>
    <footer class="footer"><span>GitHub Trophy Lab</span><span><a href="/">Studio</a> · <a href="/docs">Docs</a> · <a href="/reference">Reference</a></span></footer>
  </main>
</body>
</html>`

const referencePageHTML = `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta name="theme-color" content="#07111d">
  <title>GitHub Trophy Lab — Reference</title>
  <style>
    :root { color-scheme: dark; --ink: #f4f8fb; --muted: #9baebe; --quiet: #718294; --line: rgba(214, 237, 246, .14); --mint: #8df0cb; --radius-xl: 28px; --shadow: 0 26px 80px rgba(0,0,0,.34); }
    * { box-sizing: border-box; }
    html { background: #07111d; }
    body { min-width: 320px; min-height: 100vh; margin: 0; color: var(--ink); font-family: Inter, ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif; background: radial-gradient(circle at 10% 6%, rgba(42,151,191,.30), transparent 30rem), radial-gradient(circle at 92% 14%, rgba(77,207,168,.22), transparent 26rem), linear-gradient(145deg, #07111d 0%, #0b1826 52%, #07131c 100%); }
    .shell { width: min(1180px, calc(100% - 40px)); margin: 0 auto; padding: 60px 0 80px; }
    .topbar { display: flex; align-items: center; justify-content: space-between; padding-bottom: 40px; }
    .brand { display: inline-flex; align-items: center; gap: 12px; text-decoration: none; color: inherit; }
    .brand-mark { display: grid; width: 38px; height: 38px; place-items: center; border: 1px solid rgba(141,240,203,.42); border-radius: 12px; color: var(--mint); background: rgba(141,240,203,.1); }
    .brand-name { font-size: 14px; font-weight: 700; }
    .brand-sub { display: block; margin-top: 3px; color: var(--quiet); font-size: 11px; letter-spacing: .11em; text-transform: uppercase; }
    h1 { font-size: clamp(36px, 5vw, 64px); line-height: .96; letter-spacing: -.065em; margin: 0 0 20px; }
    .eyebrow { display: flex; align-items: center; gap: 10px; margin: 0 0 18px; color: var(--mint); font-size: 12px; font-weight: 700; letter-spacing: .16em; text-transform: uppercase; }
    .eyebrow::before { width: 28px; height: 1px; content: ""; background: var(--mint); }
    .content { line-height: 1.75; color: var(--muted); font-size: 16px; }
    .content h2 { color: var(--ink); font-size: 1.4em; margin: 2em 0 .5em; }
    .content h3 { color: var(--ink); font-size: 1.15em; margin: 1.5em 0 .4em; }
    .content code { background: rgba(255,255,255,.06); padding: 2px 7px; border-radius: 6px; font-size: .9em; color: var(--mint); }
    .content pre { background: rgba(3,13,24,.6); padding: 16px 20px; border-radius: 14px; overflow-x: auto; border: 1px solid var(--line); }
    .content pre code { background: none; padding: 0; color: var(--ink); }
    .content a { color: var(--mint); text-underline-offset: 3px; }
    .card { border: 1px solid var(--line); border-radius: var(--radius-xl); padding: 24px; margin: 20px 0; background: rgba(255,255,255,.02); }
    .card h3 { margin-top: 0; }
    .badge { display: inline-block; padding: 2px 9px; border-radius: 999px; background: rgba(141,240,203,.1); color: var(--mint); font-size: 12px; font-weight: 650; }
    .footer { padding-top: 40px; border-top: 1px solid var(--line); color: var(--quiet); font-size: 12px; display: flex; justify-content: space-between; flex-wrap: wrap; gap: 12px; }
    .footer a { color: var(--muted); }
    .footer a:hover { color: var(--mint); }
    @media (prefers-reduced-motion: reduce) { *, *::before, *::after { scroll-behavior: auto !important; transition-duration: .01ms !important; animation-duration: .01ms !important; } }
  </style>
</head>
<body>
  <main class="shell">
    <header class="topbar">
      <a class="brand" href="/" aria-label="GitHub Trophy Lab home">
        <span class="brand-mark"><svg width="21" height="21" viewBox="0 0 24 24" fill="none"><path d="M7 4h10v4.5a5 5 0 0 1-10 0V4Z" stroke="currentColor" stroke-width="1.7"/><path d="M7 6H4.8A1.8 1.8 0 0 0 3 7.8v.4a3.4 3.4 0 0 0 3.4 3.4M17 6h2.2A1.8 1.8 0 0 1 21 7.8v.4a3.4 3.4 0 0 1-3.4 3.4M12 13.5V18M8.5 21h7M9 18h6" stroke="currentColor" stroke-linecap="round" stroke-width="1.7"/></svg></span>
        <span><span class="brand-name">GitHub Trophy Lab</span><span class="brand-sub">Profile card studio</span></span>
      </a>
      <span class="badge">Reference</span>
    </header>
    <div class="content">
      <p class="eyebrow">Attribution</p>
      <h1>Built on upstream work</h1>
      <p>GitHub Trophy Lab is a self-hosted, polished wrapper around <a href="https://github.com/soulteary/github-profile-trophy" target="_blank" rel="noreferrer">soulteary/github-profile-trophy</a>. All trophy rendering logic is inherited from the original project; this service adds a studio UI, Vercel deployment, and advanced parameter support.</p>
      <div class="card">
        <h3>Original project</h3>
        <p><a href="https://github.com/soulteary/github-profile-trophy" target="_blank" rel="noreferrer">github.com/soulteary/github-profile-trophy</a></p>
        <p>MIT License · Copyright (c) 2022 soulteary</p>
      </div>
      <h2>What changed</h2>
      <ul>
        <li><strong>Studio UI</strong> — live preview, Markdown copy, advanced settings panel</li>
        <li><strong>Vercel deployment</strong> — one-click deploy with environment variable support</li>
        <li><strong>Advanced parameters</strong> — <code>rank</code>, <code>row</code>, <code>margin-h</code>, <code>title</code> with exclude/includ e support</li>
        <li><strong>Error handling</strong> — friendly messages for missing tokens and invalid usernames</li>
      </ul>
      <h2>Parameter reference</h2>
      <div class="card">
        <h3>All supported query parameters</h3>
        <p style="margin-bottom:12px">These match the original upstream API. Prefix with <code>-</code> to exclude (e.g. <code>title=-Stars,-Forks</code>).</p>
        <table style="width:100%;border-collapse:collapse;">
          <thead><tr style="border-bottom:1px solid var(--line);text-align:left"><th style="padding:8px 12px;color:var(--ink)">Param</th><th style="padding:8px 12px;color:var(--ink)">Type</th><th style="padding:8px 12px;color:var(--ink)">Default</th><th style="padding:8px 12px;color:var(--ink)">Notes</th></tr></thead>
          <tbody>
            <tr style="border-bottom:1px solid var(--line)"><td style="padding:8px 12px"><code>username</code></td><td style="padding:8px 12px">string</td><td style="padding:8px 12px">—</td><td style="padding:8px 12px">Required. GitHub username.</td></tr>
            <tr style="border-bottom:1px solid var(--line)"><td style="padding:8px 12px"><code>theme</code></td><td style="padding:8px 12px">string</td><td style="padding:8px 12px">default</td><td style="padding:8px 12px">Themes: default, onedark, dracula, nord, gruvbox, radical, tokyonight, aura</td></tr>
            <tr style="border-bottom:1px solid var(--line)"><td style="padding:8px 12px"><code>column</code></td><td style="padding:8px 12px">int</td><td style="padding:8px 12px">-1</td><td style="padding:8px 12px">-1 = auto-fit to width</td></tr>
            <tr style="border-bottom:1px solid var(--line)"><td style="padding:8px 12px"><code>row</code></td><td style="padding:8px 12px">int</td><td style="padding:8px 12px">-1</td><td style="padding:8px 12px">-1 = auto-fit to height</td></tr>
            <tr style="border-bottom:1px solid var(--line)"><td style="padding:8px 12px"><code>margin-w</code></td><td style="padding:8px 12px">int</td><td style="padding:8px 12px">25</td><td style="padding:8px 12px">Horizontal gap in px</td></tr>
            <tr style="border-bottom:1px solid var(--line)"><td style="padding:8px 12px"><code>margin-h</code></td><td style="padding:8px 12px">int</td><td style="padding:8px 12px">25</td><td style="padding:8px 12px">Vertical gap in px</td></tr>
            <tr style="border-bottom:1px solid var(--line)"><td style="padding:8px 12px"><code>no-bg</code></td><td style="padding:8px 12px">bool</td><td style="padding:8px 12px">true</td><td style="padding:8px 12px">Transparent background</td></tr>
            <tr style="border-bottom:1px solid var(--line)"><td style="padding:8px 12px"><code>no-frame</code></td><td style="padding:8px 12px">bool</td><td style="padding:8px 12px">false</td><td style="padding:8px 12px">Hide trophy border frames</td></tr>
            <tr style="border-bottom:1px solid var(--line)"><td style="padding:8px 12px"><code>title</code></td><td style="padding:8px 12px">string</td><td style="padding:8px 12px">—</td><td style="padding:8px 12px">Comma-separated titles; prefix with <code>-</code> to exclude</td></tr>
            <tr><td style="padding:8px 12px"><code>rank</code></td><td style="padding:8px 12px">string</td><td style="padding:8px 12px">—</td><td style="padding:8px 12px">Comma-separated ranks: SECRET, SSS, SS, S, AAA, AA, A, B, C; prefix with <code>-</code> to exclude</td></tr>
          </tbody>
        </table>
      </div>
      <h2>Credits</h2>
      <p>Trophy rendering engine: <a href="https://github.com/soulteary/github-profile-trophy" target="_blank" rel="noreferrer">soulteary/github-profile-trophy</a> — MIT License.</p>
      <p>UI design and Vercel integration: <a href="https://github.com/ege-bulut" target="_blank" rel="noreferrer">ege-bulut</a>.</p>
    </div>
    <footer class="footer"><span>GitHub Trophy Lab</span><span><a href="/">Studio</a> · <a href="/docs">Docs</a> · <a href="/reference">Reference</a></span></footer>
  </main>
</body>
</html>`


func renderLandingPage(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(landingPageHTML))
}

func DocsHandler(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(docsPageHTML))
}

func ReferenceHandler(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "https://github.com/ryo-ma/github-profile-trophy")
}
