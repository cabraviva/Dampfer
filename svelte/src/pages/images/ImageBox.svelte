<script lang="ts">
  import Fa from "svelte-fa";
  import "../../app.scss";
  import "./image-box.scss";
  import type { AlertType } from "../../types";
  import {
    deleteImage,
    inspectImage,
    type ListedImagesGroupedItem,
  } from "../../script/images-networks-volumes";
  import { onMount } from "svelte";
  import { getAverageBackgroundColor, searchIcons } from "../../script/icongen";
  import knorry, { del } from "knorry";
  import { getCredentials } from "../../script/login";
  import {
    Button,
    Dropdown,
    DropdownDivider,
    Modal,
    Radio,
  } from "flowbite-svelte";
  import {
    ChevronDownOutline,
    ExclamationCircleOutline,
  } from "flowbite-svelte-icons";
  import {
    faEye,
    faPen,
    faTrash,
    faXmark,
  } from "@fortawesome/free-solid-svg-icons";
  import SmallPopup from "../../popups/SmallPopup.svelte";

  let { pushAlert, updatePage, imagesWithSameTag, refetchImages } =
    $props() as {
      pushAlert: (alert: AlertType) => void;
      updatePage: (name: string) => void;
      refetchImages: () => void;
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

  // Popup states
  let showMoreInfoPopup = $state(false);
  let showDeleteModal = $state(false);
</script>

<!-- TODO:
        - Pull images
        - Authentication when pulling
        - Editing image: change image via A) Search B) Upload
        - Better image info on list page
        - Schedule re-pulling
        - Add in-use marker
        - Import & Export images
-->

<!-- View more popup -->
<SmallPopup
  show={showMoreInfoPopup}
  onclose={() => (showMoreInfoPopup = false)}
  title="View image: {imagesWithSameTag.images[0].Repository}"
>
  {#each imagesWithSameTag.images as current_image}
    {#await inspectImage(current_image.ID) then details}
      <h2>
        {current_image.Repository}<span class="darker-span"
          >:{current_image.Tag}</span
        >
      </h2>
      {#each Object.entries(details) as [cd1k, cd1v]}
        <span class="darker-span">{cd1k}:</span>
        <span class="primary-span">
          {#if Array.isArray(cd1v)}
            {JSON.stringify(cd1v)}
          {:else if typeof cd1v === "object"}
            {#each Object.entries(cd1v) as [cd2k, cd2v]}
              <br />
              <span class="darker-span ml-4">{cd2k}:</span>
              <span class="primary-span">
                {#if Array.isArray(cd2v)}
                  {JSON.stringify(cd2v)}
                {:else if typeof cd2v === "object"}
                  {JSON.stringify(cd2v)}
                {:else}
                  {cd2v}
                {/if}
              </span>
            {/each}
          {:else}
            {cd1v}
          {/if}
        </span>
        <br />
      {/each}
    {/await}
    <br /><br />
  {/each}
</SmallPopup>

<!-- Confirm delete popup -->
<Modal bind:open={showDeleteModal} size="xs" autoclose>
  <div class="text-center">
    <ExclamationCircleOutline
      class="mx-auto mb-4 text-gray-400 w-12 h-12 dark:text-gray-200"
    />
    <h3 class="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">
      Are you sure you want to delete the image {imagesWithSameTag.images.filter(
        (i) => i.Tag === selectedTag
      )[0].Repository}:{imagesWithSameTag.images.filter(
        (i) => i.Tag === selectedTag
      )[0].Tag}
    </h3>
    <Button
      color="primary"
      class="me-2"
      on:click={async () => {
        const success = await deleteImage(
          imagesWithSameTag.images[0].Repository,
          selectedTag
        );

        if (success === true) {
          pushAlert({
            icon: faTrash,
            color: "green",
            content: `Image ${imagesWithSameTag.images.filter((i) => i.Tag === selectedTag)[0].Repository}:${imagesWithSameTag.images.filter((i) => i.Tag === selectedTag)[0].Tag} was deleted successfully`,
          });
          refetchImages();
        } else {
          pushAlert({
            icon: faXmark,
            color: "red",
            content: `Error while trying to delete image ${imagesWithSameTag.images.filter((i) => i.Tag === selectedTag)[0].Repository}:${imagesWithSameTag.images.filter((i) => i.Tag === selectedTag)[0].Tag}. You might want to check the logs!`,
          });
        }
      }}>Yes, I'm sure</Button
    >
    <Button color="alternative">No, cancel</Button>
  </div>
</Modal>

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
  <div class="buttons">
    <button
      onclick={async () => {
        showMoreInfoPopup = true;
      }}
    >
      <Fa icon={faEye} />
    </button>
    <button>
      <Fa icon={faPen} />
    </button>
    <button
      class="rm"
      onclick={async () => {
        showDeleteModal = true;
      }}
    >
      <Fa icon={faTrash} />
    </button>
  </div>
</div>
