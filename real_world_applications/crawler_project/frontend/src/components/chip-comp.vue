<template>
    <div :style="{ backgroundColor: backgroundColor }"
        class="bounce w-[60vw] h-10 md:w-[15vw] text-white font-sans text-md flex flex-col items-center justify-center text-center font-medium rounded-md">
        {{ text }}
    </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
    props: {
        backgroundColor: {
            type: String,
            required: true,
        },
        text: {
            type: String,
            required: true,
        },
    },
    setup(props) {
        // Props are already defined and reactive
        return { ...props };
    },
});
</script>

<style scoped>
@keyframes horizontalBounce {

    0%,
    100% {
        transform: translateX(0);
    }

    50% {
        transform: translateX(20px);
    }
}

.bounce {
    animation: horizontalBounce 1s ease-in-out;
}
</style>
