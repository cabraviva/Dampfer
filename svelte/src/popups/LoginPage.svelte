<script lang="ts">
  import "../app.css";
  import "./LoginPage.scss";

  import {
    credentialsSaved,
    isJWTValid,
    getCredentials,
    setCredentials,
    getSavedUsername,
    setSavedUsername,
  } from "../script/login";

  import { onMount } from "svelte";
  let container: HTMLElement;

  import { Particle } from "jparticles";
  import { Input, Label } from "flowbite-svelte";
  import logo from "../../public/logo.svg";

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
</script>

<main id="animated-bg-login" class="login-page-container"></main>

<div class="login-formular">
  <div class="login-form">
    <img class=" mb-5 login-logo" src={logo} alt="Dampfer logo" />

    <h1 class="text-3xl mb-5">Login</h1>

    <div class="mb-5">
      <Label for="first_name" class="mb-2">Username</Label>
      <Input type="text" id="username" placeholder="admin" required />
    </div>

    <div class="mb-5">
      <Label for="password" class="mb-2">Password</Label>
      <Input type="password" id="password" placeholder="•••••••••" required />
    </div>
  </div>
</div>
