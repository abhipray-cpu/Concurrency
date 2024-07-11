<template>
  <!--this is the parent container-->
  <div class="w-full min-h-[100vh] flex flex-col items-center bg-gray-100 pt-10 md:pt-10">
    <!--this is the header section-->
    <div class="flex flex-col items-center">
      <h2 class="text-gray-800 font-mono text-xl md:text-3xl tracking-wider font-bold">
        Go Crawler
      </h2>
      <Vue3Lottie :animationData="CrawlerJSON" :height="computedSize" :width="computedSize" :loop="false"
        :autoplay="true" :speed="1" />
    </div>
    <!--these are the error sections-->
    <div v-if="Chip != 0" class="mb-0 md:mb-5">
      <ChipComp backgroundColor="#e57373" text="Please fill all the fields" v-if="Chip == 1" />
      <ChipComp backgroundColor="#bdbdbd" text="Please fill all the fields" v-if="Chip == 2" />
      <ChipComp backgroundColor="#ff7043" text="Something went wrong at our end" v-if="Chip == 3" />
      <ChipComp backgroundColor="#ffb74d" text="Bad values" v-if="Chip == 4" />
      <ChipComp backgroundColor="#4fc3f7" text="Account Created" v-if="Chip == 5" />
    </div>
    <!--this is the form section-->
    <div
      class="flex flex-col items-center w-full min-h-[60vh] md:w-[30vw] md:border-2 md:border-gray-400 md:rounded-md pt-10 md:pt-12">
      <input type="text" placeholder="email" v-model="email"
        class="w-[90vw] md:w-[22vw] h-12 md:h-14 mb-5 md:mb-9 pl-3 md:pl-4 font-mono font-normal text-lg md:text-xl text-gray-600 rounded-xl focus:outline-none focus:border focus:border-gray-400" />
      <input type="password" placeholder="password" v-model="password"
        class="w-[90vw] md:w-[22vw] h-12 md:h-14 mb-5 md:mb-9 pl-3 md:pl-4 font-mono font-normal text-lg md:text-xl text-gray-600 rounded-xl focus:outline-none focus:border focus:border-gray-400" />
      <div
        class="bg-gray-600 text-gray-50 font-mono text-xl md:text-2xl font-semibold flex flex-col items-center justify-center w-[60vw] md:w-[10vw] h-[12vw] md:h-10 rounded-md cursor-pointer"
        @click="login">
        Login
      </div>
      <Vue3Lottie :animationData="LoaderJSON" :height="200" :width="200" :loop="true" :autoplay="true" :speed="1"
        v-if="Loader" />
      <!--this is the alternate suignup page-->
      <span class="mt-16 md:mt-10 text-gray-700 font-sans text-lg md:text-xl" v-else>Already Have an account?
        <strong class="cursor-pointer" @click="redirect"> Login</strong>
      </span>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { onMounted, computed, ref } from "vue";
import { useCrawlerStore } from "../store";
import { useRouter } from "vue-router";
import { Vue3Lottie } from "vue3-lottie";
import ChipComp from "../components/chip-comp.vue";
import { useMediaQuery } from "@vueuse/core";
import CrawlerJSON from "../assets/crawler.json";
import LoaderJSON from "../assets/loader.json";
export default defineComponent({
  components: {
    Vue3Lottie,
    ChipComp,
  },
  setup() {
    const store = useCrawlerStore();
    const router = useRouter();
    var email = ref("");
    var password = ref("");
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
    onMounted(async () => {
      const token = store.checkAuth();
      if (token) {
        router.push({ name: "home" });
      }
    });
    const redirect = () => {
      router.push({ name: "signup" });
    };
    const login = async () => {
      if (email.value === "" || password.value === "") {
        chip.value = 1;
        setTimeout(() => {
          chip.value = 0;
        }, 1500);
        return;
      }
      const Email = email.value;
      const Password = password.value;
      loader.value = true;
      const response = await store.Login({
        email: Email,
        password: Password,
      });
      loader.value = false;
      if (response.status == 200) {
        chip.value = 5;
        setTimeout(() => {
          chip.value = 0;
          router.push({ name: "home" });
        }, 1500);
      } else if (response.status === 500) {
        chip.value = 3;
        setTimeout(() => {
          chip.value = 0;
        }, 1500);
        return;
      } else {
        chip.value = 4;
        setTimeout(() => {
          chip.value = 0;
        }, 1500);
        return;
      }
    };
    return {
      CrawlerJSON,
      computedSize,
      email,
      password,
      redirect,
      login,
      Chip,
      Loader,
      LoaderJSON,
    };
  },
});
</script>

<style scoped></style>
