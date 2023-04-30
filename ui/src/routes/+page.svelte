<script lang="ts">
	import { onMount } from "svelte"

  type Quote = {
    content?: string;
    author?: string;
  };

  async function loadQuote(): Promise<Quote> {
    const res = await fetch("/api/quotes")
    const quote = await res.json()
    if (res.ok) {
      return quote
    } else {
      throw new Error(quote)
    }
  }

  $: quote = loadQuote()
</script>

<h1>Welcome to my sandbox!</h1>

<div class="container mx-auto w-96">
  <button on:click={() => quote = loadQuote()} class="btn btn-accent btn-block btn-outline">Get a famous quote</button>

  <div class="card bg-base-100 shadow-xl">
    <div class="card-body">
      <h2 class="card-title">Famous quote</h2>
      {#await quote then quote}
          <p>{quote.content}</p>
          <p>{quote.author}</p>
      {:catch error}
          <p>Error: {JSON.stringify(error)}</p>
      {/await}
    </div>
  </div>
</div>

<svelte:head>
  <title>Homepage</title>
</svelte:head>
