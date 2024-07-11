<template>
    <div class="w-100 h-100 flex flex-col items-center justify-center">
        <Vue3Lottie :animationData="ErrorJSON" :height="computedSize" :width="computedSize" :loop="false"
            :autoplay="true" :speed="1" />

        <div class="bg-gray-600 text-gray-50 font-mono text-xl md:text-2xl font-semibold flex flex-col items-center justify-center w-[60vw] md:w-[10vw] h-[12vw] md:h-10 rounded-md cursor-pointer"
            @click="redirect">
            Home
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent, computed } from "vue";
import { Vue3Lottie } from "vue3-lottie";
import { useMediaQuery } from "@vueuse/core";
import ErrorJSON from "../assets/404.json";
import { useRouter } from "vue-router";
export default defineComponent({
    components: { Vue3Lottie },
    setup() {
        var isMobile = useMediaQuery("(max-width: 480px)");
        const isTablet = useMediaQuery("(min-width: 768px)");
        const computedSize = computed(() => {
            if (isMobile) return 340;
            if (isTablet) return 400;
            return 550;
        });
        const router = useRouter();
        const redirect = () => {
            router.push({ name: "home" });
        };
        return {
            ErrorJSON,
            computedSize,
            redirect,
        };
    },
});
</script>
