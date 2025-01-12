<script lang="ts">
  import "../app.scss";

  import { credentialsSaved, isJWTValid } from "../script/login";
  import LoginPage from "./LoginPage.svelte";
  import NeedsPwChangeCheck from "./NeedsPWChangeCheck.svelte";

  const { pushAlert, updatePage } = $props();
</script>

<main>
  {#if credentialsSaved()}
    {#await isJWTValid() then isValid}
      {#if isValid}
        <!-- Logged in -->
        <NeedsPwChangeCheck {pushAlert} {updatePage} />
      {:else}
        <LoginPage />
      {/if}
    {:catch error}
      <p style="color: red">{error}</p>
    {/await}
  {:else}
    <LoginPage />
  {/if}
</main>
