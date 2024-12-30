<script lang="ts">
  import "../app.css";
  import "./LoginPage.scss";

  import {
    credentialsSaved,
    isJWTValid,
    getCredentials,
    setCredentials,
    login,
    getSavedUsername,
    setSavedUsername,
  } from "../script/login";

  import { onMount } from "svelte";
  let container: HTMLElement;

  import { Particle } from "jparticles";
  import { Input, Label } from "flowbite-svelte";
  import logo from "../../public/logo.svg";
  import Fa from "svelte-fa";
  import { faArrowRight } from "@fortawesome/free-solid-svg-icons";
  import { Button } from "flowbite-svelte";
  import { Alert } from "flowbite-svelte";
  import { type GeneralAlert } from "../types";

  onMount(() => {
    new Particle("#animated-bg-login", {
      color: "#00D0E9",
      lineShape: "cube",
      range: 2000,
      proximity: 100,
      // Turn on parallax effect
      parallax: true,
    });
  });

  // Logic
  let alert: GeneralAlert = $state({ bold: "", msg: "", color: "red" });
  let username = $state(getSavedUsername() || "");
  let password = $state("");

  async function submitLogin() {
    setSavedUsername(username);
    const [loginSuccessful, bold, msg] = await login(username, password);
    if (loginSuccessful) {
      // Refresh page
      // @ts-expect-error
      window.location = window.location.href;
    } else {
      alert = {
        bold,
        msg,
        color: "red",
      };
    }
  }
</script>

<main id="animated-bg-login" class="login-page-container"></main>

<div class="login-formular">
  <div class="login-form">
    <img class=" mb-5 login-logo" src={logo} alt="Dampfer logo" />

    <h1 class="text-3xl mb-5">Login</h1>

    <div class="mb-5">
      <Label for="username" class="mb-2">Username</Label>
      <Input
        type="text"
        id="username"
        placeholder="admin"
        bind:value={username}
      />
    </div>

    <div class="mb-5">
      <Label for="password" class="mb-2">Password</Label>
      <Input
        type="password"
        id="password"
        placeholder="•••••••••"
        bind:value={password}
      />
    </div>

    <Button class="mt-2" on:click={submitLogin}
      >Submit <Fa class="ml-2" icon={faArrowRight} /></Button
    >
    {#if alert.msg !== ""}
      <Alert class="mt-5" color={alert.color}>
        <span class="font-bold">{alert.bold}</span>
        {alert.msg}
      </Alert>
    {/if}
  </div>
</div>
