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
  import { Button, Dropdown, DropdownDivider, Radio } from "flowbite-svelte";
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
  const sortedTags: (string | null)[] = [null, null];
  const preSortedTags = imagesWithSameTag.tags.sort();

  for (const tagEl of preSortedTags) {
    if (tagEl === "latest") {
      sortedTags[0] = "latest";
    } else if (tagEl === "nightly") {
      sortedTags[1] = "nightly";
    } else {
      sortedTags.push(tagEl);
    }
  }

  let firstTagNonNull: string;

  if (sortedTags[0] !== null) {
    firstTagNonNull = sortedTags[0];
  } else if (sortedTags[1] !== null) {
    firstTagNonNull = sortedTags[1];
  } else {
    firstTagNonNull = `${sortedTags[2]}`;
  }

  let selectedTag = $state(firstTagNonNull);
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
        {#if sortedTags.filter((t) => t !== null).length > 1}
          <ChevronDownOutline class="w-6 h-6 ms-2 text-white dark:text-white" />
        {/if}
      </Button>
      {#if sortedTags.filter((t) => t !== null).length > 1}
        <Dropdown class="w-48 p-3 space-y-1">
          {#if sortedTags[0] !== null || sortedTags[1] !== null}
            {#each sortedTags.slice(0, 2) as tag}
              {#if tag !== null}
                <li
                  class="rounded p-2 hover:bg-gray-100 dark:hover:bg-gray-600"
                >
                  <Radio name="selectedTag" bind:group={selectedTag} value={tag}
                    >{tag}</Radio
                  >
                </li>
              {/if}
            {/each}
            <DropdownDivider />
          {/if}
          {#each sortedTags.slice(2) as tag}
            <li class="rounded p-2 hover:bg-gray-100 dark:hover:bg-gray-600">
              <Radio name="selectedTag" bind:group={selectedTag} value={tag}
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
