<script lang="ts">
  import "../app.scss";

  import { credentialsSaved, isJWTValid } from "../script/login";
  import LoginPage from "./LoginPage.svelte";
</script>

<main>
  {#if credentialsSaved()}
    {#await isJWTValid() then isValid}
      {#if isValid}
        <!-- Loggedin -->
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
