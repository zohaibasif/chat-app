<script>
  // imports
  import { onMount } from 'svelte';
  import Pusher from 'pusher-js';

  // variables
  let username = "username";
  let message = "";
  let messages = [];

  // functions
  onMount(() => {
    Pusher.logToConsole = true;

    const pusher = new Pusher("bcc34fc5bfaef1df32bc", {
      cluster: "ap2",
    });

    const channel = pusher.subscribe("simple-chat-app");
    channel.bind("message", (data) => {
      messages = [...messages, data];
    });
  });

  const submit = async () => {
    await fetch('http://127.0.0.1:8000/api/messages', {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        username,
        message,
      }),
    });

	message = '';
  };
</script>

<div class="container">
  <div
    class="d-flex flex-column align-items-stretch flex-shrink-0 bg-white"
    style="width: 380px;"
  >
    <div
      class="d-flex align-items-center flex-shrink-0 p-3 link-dark text-decoration-none border-bottom"
    >
      <input class="fs-5 fw-semibold" bind:value={username} />
    </div>
    <div class="list-group list-group-flush border-bottom scrollarea">
      {#each messages as msg}
        <div class="list-group-item list-group-item-action py-3 lh-sm">
          <div class="d-flex w-100 align-items-center justify-content-between">
            <strong class="mb-1">{msg.username}</strong>
          </div>
          <div class="col-10 mb-1 small">
            {msg.message}
          </div>
        </div>
      {/each}
    </div>
  </div>
  <form on:submit|preventDefault={submit}>
    <input
      class="form-control"
      placeholder="#Write a message"
      bind:value={message}
    />
  </form>
</div>

<style>
  .scrollarea {
    min-height: 500px;
  }
</style>
