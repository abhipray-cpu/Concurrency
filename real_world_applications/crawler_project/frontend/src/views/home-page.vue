<template>
  <div class="flex flex-col items-center w-full min-h-[100vh] bg-gray-100 pt-4">
    <!--head section-->
    <div class="w-full flex flex-col items-center">
      <div class="flex flex-row justify-between w-full px-4">
        <div class="mt-4">
          <h2 class="font-sans text-gray-700 font-semibold text-lg">Go Crawler</h2>
          <h3 class="font-sans text-gray-500 font-semibold text-base w-[60vw]">
            A concurrent web crawler based on Go
          </h3>
        </div>
        <Vue3Lottie :animationData="ProfileJSON" :height="computedSize" :width="computedSize" :loop="true"
          :autoplay="true" :speed="1" class="cursor-pointer" @click="redirect('profile')" />
      </div>
      <div class="flex flex-col items-center justify-center mt-4 md:-mt-28">
        <input type="text" placeholder="Search Page.."
          class="w-[90vw] md:w-[30vw] h-10 md:h-12 mb-5 md:mb-9 pl-3 md:pl-4 font-mono font-normal text-lg md:text-xl text-gray-600 rounded-xl focus:outline-none cursor-pointer"
          @click="redirect('search')" />
      </div>
    </div>
    <!--loading-->
    <Vue3Lottie :animationData="LoaderJSON" :height="300" :width="300" :loop="true" :autoplay="true" :speed="1"
      class="cursor-pointer" v-if="Loader" />
    <!--chips-->
    <!--these are the error sections-->
    <div v-if="Chip != 0" class="mb-0 md:mb-5">
      <ChipComp backgroundColor="#e57373" text="Something went wrong at our end" v-if="Chip == 1" />
      <ChipComp backgroundColor="#bdbdbd" text="Failed to fetch pages" v-if="Chip == 2" />
      <ChipComp backgroundColor="#ff7043" text="Something went wrong at our end" v-if="Chip == 3" />
      <ChipComp backgroundColor="#ffb74d" text="Bad values" v-if="Chip == 4" />
      <ChipComp backgroundColor="#4fc3f7" text="Account Created" v-if="Chip == 5" />
    </div>
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
import { defineComponent, computed, ref } from "vue";
import { onMounted } from "vue";
import { useRouter } from "vue-router";
import { useCrawlerStore } from "../store";
import { Vue3Lottie } from "vue3-lottie";
import ChipComp from "../components/chip-comp.vue";
import { useMediaQuery } from "@vueuse/core";
import ProfileJSON from "../assets/Profile.json";
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
    const route = useRouter();
    const store = useCrawlerStore();
    var isMobile = useMediaQuery("(max-width: 480px)");
    const isTablet = useMediaQuery("(min-width: 768px)");
    const computedSize = computed(() => {
      if (isMobile) return 140;
      if (isTablet) return 300;
      return 350;
    });
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
    onMounted(async () => {
      const auth = store.checkAuth();
      if (!auth) {
        route.push("/login");
      }
      loader.value = true;
      const response = await store.GetPages();
      loader.value = false;
      if (response.status === 200) {
        data.value = response.data;
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
        return;
      }
    });
    const redirect = (val: string) => {
      route.push({ name: val });
    };

    const openpage = (id: string) => {
      route.push({ name: "page", params: { id: id } });
    };
    return {
      ProfileJSON,
      computedSize,
      Loader,
      LoaderJSON,
      Chip,
      redirect,
      Data,
      openpage,
    };
  },
});
</script>

<style scoped></style>
