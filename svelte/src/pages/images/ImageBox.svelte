<script lang="ts">
  import Fa from "svelte-fa";
  import "../../app.scss";
  import "./image-box.scss";
  import type { AlertType } from "../../types";
  import { type ListedImagesGroupedItem } from "../../script/images-networks-volumes";
  import { onMount } from "svelte";
  import { getAverageBackgroundColor, searchIcons } from "../../script/icongen";
  import knorry from "knorry";
  import { getCredentials } from "../../script/login";
  import { Button, Dropdown, Radio } from "flowbite-svelte";
  import { ChevronDownOutline } from "flowbite-svelte-icons";

  let { pushAlert, updatePage, imagesWithSameTag } = $props() as {
    pushAlert: (alert: AlertType) => void;
    updatePage: (name: string) => void;
    imagesWithSameTag: ListedImagesGroupedItem;
  };

  let iconContainer: HTMLDivElement | null = $state(null);

  onMount(async () => {
    async function calcBg() {
      try {
        if (iconContainer !== null) {
          const avgBgColor = await getAverageBackgroundColor(
            iconContainer.querySelector("img") || document.createElement("img")
          );

          iconContainer.style.backgroundColor = avgBgColor;
        } else {
          setTimeout(calcBg, 50);
        }
      } catch {
        setTimeout(calcBg, 50);
      }
    }
    calcBg();
  });

  // TODO: sort tags so that latest is on top, 2. is nightly and then all others are sorted by default array.sort

  let selectedTag = $state(imagesWithSameTag.tags[0]);
</script>

<!-- TODO:
        - Pull images
        - Authentication when pulling
        - Editing image: Delete, change image via A) Search B) Upload
        - Better image info on list page
        - Detailed image info when clicking on image (using image inspect)
        - Group images by tag
        - Schedule re-pulling
        - Add in-use marker
        - Add prompts if deleting when in use
-->

<div class="image-box">
  <div class="icon" bind:this={iconContainer}>
    {#await fetch( `/api/icongen/get-icon?id=${encodeURIComponent(imagesWithSameTag.images[0].ID)}`, { headers: { Authorization: `Bearer ${getCredentials()}` } } ) then imgRes}
      {#await imgRes.blob() then blob}
        <img src={URL.createObjectURL(blob)} alt="" />
      {/await}
    {/await}
  </div>
  <div class="info">
    <span class="repo">{imagesWithSameTag.images[0].Repository}</span>
    <span class="tags">
      <Button>
        :{selectedTag}
        {#if imagesWithSameTag.tags.length > 1}
          <ChevronDownOutline class="w-6 h-6 ms-2 text-white dark:text-white" />
        {/if}
      </Button>
      {#if imagesWithSameTag.tags.length > 1}
        <Dropdown class="w-48 p-3 space-y-1">
          {#each imagesWithSameTag.tags as tag}
            <li class="rounded p-2 hover:bg-gray-100 dark:hover:bg-gray-600">
              <Radio name="group2" bind:group={selectedTag} value={tag}
                >{tag}</Radio
              >
            </li>
          {/each}
        </Dropdown>
      {/if}
    </span>
  </div>
  <div class="buttons"></div>
</div>
