<template>
    <div class="flex flex-col items-center w-full min-h-[100vh] bg-gray-100 pt-4">
        <!--these are the error sections-->
        <div v-if="Chip != 0" class="mb-0 md:mb-5">
            <ChipComp backgroundColor="#e57373" text="Something went wrong at our end" v-if="Chip == 1" />
            <ChipComp backgroundColor="#bdbdbd" text="Failed to fetch page details" v-if="Chip == 2" />
            <ChipComp backgroundColor="#ff7043" text="Something went wrong at our end" v-if="Chip == 3" />
            <ChipComp backgroundColor="#ffb74d" text="Deleted the page" v-if="Chip == 4" />
            <ChipComp backgroundColor="#4fc3f7" text="Please add some content" v-if="Chip == 5" />
            <ChipComp backgroundColor="#4fc3f7" text="Content updated" v-if="Chip == 6" />
        </div>
        <!--loader-->
        <Vue3Lottie :animationData="LoaderJSON" :height="300" :width="300" :loop="true" :autoplay="true" :speed="1"
            class="cursor-pointer" v-if="Loader" />
        <!--page details-->
        <div class="flex flex-col justify-center mt-5 md:mt-6" v-if="Current === 0">
            <h3 class="font-sans text-lg md:text-xl text-center font-bold tracking-wider text-gray-700">
                Page Details
            </h3>

            <div class="flex flex-row gap-4 w-full justify-end mt-5">
                <div @click="togglePage(1)"
                    class="w-[30vw] h-8 md:w-[10vw] bg-blue-500 flex flex-col items-center justify-center text-gray-50 font-sans text-base md:text-lg rounded font-medium cursor-pointer">
                    Edit
                </div>
                <div class="w-[30vw] h-8 md:w-[10vw] bg-red-500 flex flex-col items-center justify-center text-gray-50 font-sans text-base md:text-lg mr-4 rounded font-medium cursor-pointer"
                    @click="togglePage(2)">
                    Delete
                </div>
            </div>
            <div class="flex flex-col justify-center gap-4 ml-4 mt-4">
                <div class="">
                    <h2 class="text-gray-800 font-sans font-bold tracking-wide text-xl md:text-2xl">
                        Title:
                    </h2>
                    <h2>{{ Title }}</h2>
                </div>
                <div class="">
                    <h2 class="text-gray-800 font-sans font-bold tracking-wide text-xl md:text-2xl">
                        URL:
                    </h2>
                    <h2>{{ Url }}</h2>
                </div>
                <div class="">
                    <h2 class="text-gray-800 font-sans font-bold tracking-wide text-xl md:text-2xl">
                        Description:
                    </h2>
                    <h2>{{ Description }}</h2>
                </div>
                <div class="">
                    <h2 class="text-gray-800 font-sans font-bold tracking-wide text-xl md:text-2xl">
                        Keywords:
                    </h2>
                    <h2>{{ Keywords }}</h2>
                </div>
                <div class="">
                    <h2 class="text-gray-800 font-sans font-bold tracking-wide text-xl md:text-2xl">
                        Content:
                    </h2>
                    <h2>{{ Content }}</h2>
                </div>
            </div>
        </div>
        <div class="flex flex-col justify-center mt-5 md:mt-6" v-if="Current === 1">
            <h3 class="text-center font-bold font-mono text-xl md:text-2xl text-gray-700">
                Add new content
            </h3>
            <textarea name="" id="" cols="10" rows="15" class="w-[90vw] md:w-[60vw]" v-model="newContent"></textarea>
            <div class="flex flex-tow items-center justify-center gap-5 mt-10">
                <div @click="togglePage(0)"
                    class="w-[30vw] h-8 md:w-[10vw] bg-blue-500 flex flex-col items-center justify-center text-gray-50 font-sans text-base md:text-lg rounded font-medium cursor-pointer">
                    Cancel
                </div>
                <div @click="updateFunc"
                    class="w-[30vw] h-8 md:w-[10vw] bg-red-500 flex flex-col items-center justify-center text-gray-50 font-sans text-base md:text-lg mr-4 rounded font-medium cursor-pointer">
                    Update
                </div>
            </div>
        </div>
        <div class="flex flex-col justify-center mt-20 md:mt-26" v-if="Current === 2">
            <h3 class="text-center font-bold font-mono text-xl md:text-2xl text-gray-700">
                Are you sure you want to delete this page?
            </h3>
            <div class="flex flex-tow items-center justify-center gap-5 mt-10">
                <div @click="togglePage(0)"
                    class="w-[30vw] h-8 md:w-[10vw] bg-blue-500 flex flex-col items-center justify-center text-gray-50 font-sans text-base md:text-lg rounded font-medium cursor-pointer">
                    Cancel
                </div>
                <div @click="delFunc"
                    class="w-[30vw] h-8 md:w-[10vw] bg-red-500 flex flex-col items-center justify-center text-gray-50 font-sans text-base md:text-lg mr-4 rounded font-medium cursor-pointer">
                    Delete
                </div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref, computed } from "vue";
import { useRouter, useRoute } from "vue-router";
import { useCrawlerStore } from "../store";
import ChipComp from "../components/chip-comp.vue";
import { Vue3Lottie } from "vue3-lottie";
import LoaderJSON from "../assets/loader.json";
export default defineComponent({
    components: { Vue3Lottie, ChipComp },
    setup() {
        const route = useRoute();
        const router = useRouter();
        const store = useCrawlerStore();
        var chip = ref(0);
        var content = ref("");
        var description = ref("");
        var keywords = ref([]);
        var title = ref("");
        var url = ref("");
        var PageId = ref("");
        var newContent = ref("");
        const Content = computed(() => {
            return content.value === "" ? "Not Available" : content.value;
        });
        const Description = computed(() => {
            return description.value === "" ? "Not Available" : description.value;
        });
        const Keywords = computed(() => {
            return keywords.value;
        });
        const Title = computed(() => {
            return title.value;
        });
        const Url = computed(() => {
            return url.value === "" ? "Not Available" : url.value;
        });
        const Chip = computed(() => {
            return chip.value;
        });
        const loader = ref(false);
        const Loader = computed(() => {
            return loader.value;
        });
        const current = ref(0);
        const Current = computed(() => {
            return current.value;
        });
        const togglePage = (val: number) => {
            console.log(val);
            current.value = val;
        };
        onMounted(async () => {
            const id = Array.isArray(route.params.id) ? route.params.id[0] : route.params.id;
            loader.value = true;
            const response = await store.GetPage(id);
            loader.value = false;
            if (response.status === 200) {
                PageId.value = response.data.id;
                content.value = response.data.content;
                newContent.value = response.data.content;
                description.value = response.data.description;
                keywords.value = response.data.keywords;
                title.value = response.data.title;
                url.value = response.data.url;
                return;
            } else if (response.status === 500) {
                chip.value = 1;
                setTimeout(() => {
                    chip.value = 0;
                }, 1500);
                return;
            } else {
                chip.value = 2;
                setTimeout(() => {
                    chip.value = 0;
                }, 1500);
            }
        });
        const delFunc = async () => {
            current.value = 0;
            loader.value = true;
            const response = await store.DeletePage(PageId.value);
            loader.value = false;
            if (response.status === 200) {
                chip.value = 4;
                setTimeout(() => {
                    chip.value = 0;
                }, 1500);
                router.push("/home");
                return;
            } else if (response.status === 500) {
                chip.value = 1;
                setTimeout(() => {
                    chip.value = 0;
                }, 1500);
                return;
            } else {
                chip.value = 3;
                setTimeout(() => {
                    chip.value = 0;
                }, 1500);
            }
        };
        const updateFunc = async () => {
            current.value = 0;
            if (newContent.value === "") {
                chip.value = 5;
                setTimeout(() => {
                    chip.value = 0;
                }, 1500);
                return;
            }
            loader.value = true;
            const response = await store.EditPage(PageId.value, newContent.value);
            loader.value = false;
            if (response.status === 200) {
                chip.value = 6;
                setTimeout(() => {
                    chip.value = 0;
                }, 1500);
                content.value = newContent.value;
                return;
            } else if (response.status === 500) {
                chip.value = 1;
                setTimeout(() => {
                    chip.value = 0;
                }, 1500);
                return;
            } else {
                chip.value = 3;
                setTimeout(() => {
                    chip.value = 0;
                }, 1500);
            }
        };

        return {
            Chip,
            Loader,
            LoaderJSON,
            Content,
            Description,
            Keywords,
            Title,
            Url,
            Current,
            togglePage,
            delFunc,
            updateFunc,
            newContent,
        };
    },
});
</script>

<style scoped></style>
