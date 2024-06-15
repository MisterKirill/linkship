<script setup>
import { Link } from '@inertiajs/vue3';
</script>

<template>
    <main>
        <header>
            <Link href="/">
                <img src="../../assets/linkship.svg" alt="logo" width="120">
            </Link>

            <nav>
                <Link href="/register">
                    <button>Register</button>
                </Link>
                <Link href="/login">
                    <button>Login</button>
                </Link>
            </nav>
        </header>

        <slot />
    </main>
</template>

<style scoped>
header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1rem 3rem;
}
</style>
