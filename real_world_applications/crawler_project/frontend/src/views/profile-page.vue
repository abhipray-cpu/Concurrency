<template>
    <div class="flex flex-col items-center w-full min-h-[100vh] bg-gray-100 pt-4">
        <div>
            <Vue3Lottie :animationData="ProfileJSON" :height="computedSize" :width="computedSize" :loop="true"
                :autoplay="true" :speed="1" class="cursor-pointer" />
            <h1 class="text-center font-mono text-gray-700 tracking-wider text-xl md:text-2xl font-semibold">
                {{ UserName }}
            </h1>
            <div v-if="Chip != 0" class="mb-0 md:mb-5 w-full flex flex-col justify-center">
                <ChipComp backgroundColor="#e57373" text="Something went wrong at our end" v-if="Chip == 1" />
                <ChipComp backgroundColor="#bdbdbd" text="Failed to add url" v-if="Chip == 2" />
                <ChipComp backgroundColor="#ff7043" text="URL addedd successfully" v-if="Chip == 3" />
                <ChipComp backgroundColor="#ffb74d" text="Bad values" v-if="Chip == 4" />
                <ChipComp backgroundColor="#4fc3f7" text="Account Created" v-if="Chip == 5" />
            </div>
            <Vue3Lottie :animationData="LoaderJSON" :height="300" :width="300" :loop="true" :autoplay="true" :speed="1"
                class="cursor-pointer" v-if="Loader" />
            <div class="mt-20 md:mt-6 flex flex-col items-center">
                <input type="text" placeholder="Add a URL to crawl" v-model="url"
                    class="w-[90vw] md:w-[30vw] h-11 mb-5 md:mb-9 pl-3 md:pl-4 font-mono font-normal text-lg md:text-xl text-gray-600 rounded-xl focus:outline-none cursor-pointer"
                    @click="addURL" />
                <div @click="addURL"
                    class="bg-gray-600 text-gray-50 font-mono text-xl md:text-2xl font-semibold flex flex-col items-center justify-center w-[60vw] md:w-[10vw] h-9 rounded-md cursor-pointer">
                    Add URL
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent, computed, onMounted, ref } from "vue";
import { Vue3Lottie } from "vue3-lottie";
import ProfileJSON from "../assets/Profile.json";
import LoaderJSON from "../assets/loader.json";
import { useMediaQuery } from "@vueuse/core";
import { useCrawlerStore } from "../store";
import ChipComp from "../components/chip-comp.vue";
export default defineComponent({
    components: { Vue3Lottie, ChipComp },
    setup() {
        var isMobile = useMediaQuery("(max-width: 480px)");
        const isTablet = useMediaQuery("(min-width: 768px)");
        const url = ref("");
        const computedSize = computed(() => {
            if (isMobile) return 140;
            if (isTablet) return 300;
            return;
            350;
        });
        var username = ref("");
        var UserName = computed(() => {
            return username.value;
        });
        var chip = ref(0);
        const Chip = computed(() => {
            return chip.value;
        });
        const loader = ref(false);
        const Loader = computed(() => {
            return loader.value;
        });
        const store = useCrawlerStore();
        onMounted(async () => {
            const response = await store.Profile();
            console.log(response);
            if (response.status === 200) {
                username.value = response.data.username;
            } else if (response.status === 500) {
                chip.value = 4;
                setTimeout(() => {
                    chip.value = 0;
                }, 3000);
                return;
            } else {
                chip.value = 5;
                setTimeout(() => {
                    chip.value = 0;
                }, 3000);
                return;
            }
        });
        const addURL = async () => {
            if (url.value === "") {
                return;
            }
            loader.value = true;
            url.value = "";
            const response = await store.AddURL(url.value);
            loader.value = false;
            if (response.status === 200) {
                chip.value = 3;
                setTimeout(() => {
                    chip.value = 0;
                }, 3000);
                return;
            } else if (response.status === 500) {
                chip.value = 1;
                setTimeout(() => {
                    chip.value = 0;
                }, 3000);
                return;
            } else {
                chip.value = 2;
                setTimeout(() => {
                    chip.value = 0;
                }, 3000);
                return;
            }
        };
        return {
            ProfileJSON,
            computedSize,
            UserName,
            Chip,
            Loader,
            LoaderJSON,
            url,
            addURL,
        };
    },
});
</script>

<style scoped></style>
