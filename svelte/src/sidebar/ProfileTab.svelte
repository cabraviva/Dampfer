<script lang="ts">
  import Fa from "svelte-fa";
  import { Button, Dropdown, DropdownItem } from "flowbite-svelte";

  import "../app.scss";
  import "../sidebar.scss";
  import { faUserCircle } from "@fortawesome/free-solid-svg-icons";
  import { setCredentials } from "../script/login";
  import { whoami } from "../script/whoami";
</script>

<Fa icon={faUserCircle} style="outline: none;" />

{#await whoami() then userInfo}
  <Dropdown style="outline: none;">
    <DropdownItem slot="header">Hi, {userInfo.username}!</DropdownItem>

    <DropdownItem
      slot="footer"
      on:click={() => {
        setCredentials("");
        // @ts-expect-error
        location = location.href;
      }}>Sign out</DropdownItem
    >
  </Dropdown>
{/await}
