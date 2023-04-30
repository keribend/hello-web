import { c as create_ssr_component, i as is_promise, n as noop, e as escape } from "../../chunks/index.js";
async function loadQuote() {
  const res = await fetch("/api/quotes");
  const quote2 = await res.json();
  if (res.ok) {
    return quote2;
  } else {
    throw new Error(quote2);
  }
}
const Page = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let quote;
  quote = loadQuote();
  return `<h1>Welcome to my sandbox!</h1>

<div class="container mx-auto w-96"><button class="btn btn-accent btn-block btn-outline">Get a famous quote</button>

  <div class="card bg-base-100 shadow-xl"><div class="card-body"><h2 class="card-title">Famous quote</h2>
      ${function(__value) {
    if (is_promise(__value)) {
      __value.then(null, noop);
      return ``;
    }
    return function(quote2) {
      return `
          <p>${escape(quote2.content)}</p>
          <p>${escape(quote2.author)}</p>
      `;
    }(__value);
  }(quote)}</div></div></div>

${$$result.head += `<!-- HEAD_svelte-61oa49_START -->${$$result.title = `<title>Homepage</title>`, ""}<!-- HEAD_svelte-61oa49_END -->`, ""}`;
});
export {
  Page as default
};
