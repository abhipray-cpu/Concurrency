<template>
    <div class="flex flex-col items-center w-full min-h-[100vh] bg-gray-100 pt-4">
        <input type="text" placeholder="Search Page.."
            class="w-[90vw] md:w-[30vw] h-10 md:h-12 mb-5 md:mb-9 pl-3 md:pl-4 font-mono font-normal text-lg md:text-xl text-gray-600 rounded-xl focus:outline-none cursor-pointer"
            @keyup.enter="searchFun" v-model="searchVal" />

        <!--chips-->
        <!--these are the error sections-->
        <div v-if="Chip != 0" class="mb-0 md:mb-5">
            <ChipComp backgroundColor="#e57373" text="Please enter a value" v-if="Chip == 1" />
            <ChipComp backgroundColor="#bdbdbd" text="Somethig went wrong at our end" v-if="Chip == 2" />
            <ChipComp backgroundColor="#ff7043" text="Faile to searh pages" v-if="Chip == 3" />
            <ChipComp backgroundColor="#ffb74d" text="No matching result" v-if="Chip == 4" />
            <ChipComp backgroundColor="#4fc3f7" text="Account Created" v-if="Chip == 5" />
        </div>
        <!--loader-->
        <Vue3Lottie :animationData="LoaderJSON" :height="300" :width="300" :loop="true" :autoplay="true" :speed="1"
            class="cursor-pointer" v-if="Loader" />
        <!--crawled urls-->
        <div class="flex flex-row flex-wrap justify-center items-center mt-4 gap-5 pb-5">
            <div v-for="(data, index) in Data" :key="index"
                class="bg-gray-600 rounded-md flex flex-col items-center justify-center cursor-pointer w-[90vw] md:w-[30vw] h-14 shadow-md overflow-hidden"
                @click="openpage(data.id)">
                <h2 v-if="data.title == ''" class="text-gray-50 font-sans font-medium text-xl md:text-lg text-center">
                    Untitled
                </h2>
                <h2 class="text-gray-50 font-sans font-medium text-xl md:text-lg text-center">
                    {{ data.title }}
                </h2>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted, computed } from "vue";
import { useRouter } from "vue-router";
import { useCrawlerStore } from "../store";
import ChipComp from "../components/chip-comp.vue";
import { Vue3Lottie } from "vue3-lottie";
import LoaderJSON from "../assets/loader.json";
export default defineComponent({
    components: { Vue3Lottie, ChipComp },
    setup() {
        interface Page {
            id: string;
            url: string;
            statusCode: number;
            content: string;
            crawledAt: Date;
            title: string;
            description: string;
            keywords: string[];
        }
        var searchVal = ref("");
        const route = useRouter();
        const store = useCrawlerStore();
        var chip = ref(0);
        const Chip = computed(() => {
            return chip.value;
        });
        const loader = ref(false);
        const Loader = computed(() => {
            return loader.value;
        });
        var data = ref<Page[]>([]);
        const Data = computed(() => {
            return data.value;
        });
        onMounted(() => {
            const auth = store.checkAuth();
            if (!auth) {
                route.push("/login");
            }
        });

        const searchFun = async () => {
            if (searchVal.value === "") {
                chip.value = 1;
                setTimeout(() => {
                    chip.value = 0;
                }, 1500);
                return;
            }
            loader.value = true;
            const response = await store.SearchPage(searchVal.value);
            loader.value = false;
            searchVal.value = "";
            if (response.status === 200) {
                let pages: Page[] = response.data;
                if (pages.length === 0) {
                    chip.value = 4;
                    setTimeout(() => {
                        chip.value = 0;
                    }, 1500);
                }
                data.value = pages;
                return;
            } else if (response.status === 500) {
                chip.value = 2;
                setTimeout(() => {
                    chip.value = 0;
                }, 1500);
                return;
            } else {
                chip.value = 3;
                setTimeout(() => {
                    chip.value = 0;
                }, 1500);
                return;
            }
        };
        const openpage = (id: string) => {
            route.push({ name: "page", params: { id: id } });
        };
        return {
            searchVal,
            Data,
            Chip,
            Loader,
            LoaderJSON,
            searchFun,
            openpage,
        };
    },
});
</script>
