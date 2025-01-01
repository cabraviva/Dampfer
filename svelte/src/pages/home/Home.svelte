<script lang="ts">
  import "../../app.scss";
  import logo from "../../../public/logo.svg";
  import SuggestionItem from "./SuggestionItem.svelte";
  import { faDocker } from "@fortawesome/free-brands-svg-icons";
  import {
    faBox,
    faClockRotateLeft,
    faGears,
    faPlus,
    faSdCard,
    faServer,
  } from "@fortawesome/free-solid-svg-icons";
  import { whoami } from "../../script/whoami";

  let { pushAlert, updatePage } = $props();
</script>

<main class="flex justify-center flex-col w-full h-full items-center p-2">
  <div
    class="flex justify-center flex-col items-center"
    style="margin-bottom: 5rem;"
  >
    <img src={logo} alt="Dampfer Logo" style="height: 30vh;" />
    <h1 class="text-5xl">Welcome to Dampfer</h1>
  </div>

  <div
    class="mb-8 welcome-suggestions flex flex-col justify-center items-center"
  >
    <h2>Start by...</h2>

    {#await whoami() then userInfo}
      <div class="flex flex-row justify-center items-center mb-6 mt-6">
        <!-- 1 -->
        <SuggestionItem
          icon={faDocker}
          text="Checking out your containers..."
          href={() => updatePage("containers")}
        />

        <!-- 2 -->
        {#if userInfo.admin}
          <SuggestionItem
            icon={faPlus}
            text="Creating a new service..."
            href={() => {
              updatePage("containers");
              // TODO: Open add container screen
            }}
          />
        {:else}
          <SuggestionItem
            icon={faBox}
            text="Checking out your images..."
            href={() => {
              updatePage("images");
            }}
          />
        {/if}
      </div>
      <div class="flex flex-row justify-center items-center">
        <!-- 3 -->
        {#if userInfo.admin}
          <SuggestionItem
            icon={faClockRotateLeft}
            text="Exporting backups..."
            href={() => {
              updatePage("backup");
            }}
          />
        {:else}
          <SuggestionItem
            icon={faSdCard}
            text="Checking out your volumes..."
            href={() => {
              updatePage("volumes");
            }}
          />
        {/if}

        <!-- 4 -->
        {#if userInfo.systemAdmin}
          <SuggestionItem
            icon={faGears}
            text="Configuring Dampfer..."
            href={() => {
              updatePage("settings");
            }}
          />
        {:else}
          <SuggestionItem
            icon={faServer}
            text="Checking your system stats..."
            href={() => {
              updatePage("system");
            }}
          />
        {/if}
      </div>
    {/await}
  </div>
</main>
