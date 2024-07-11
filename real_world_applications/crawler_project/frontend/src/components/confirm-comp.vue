<template>
    <div class="popup-overlay" v-if="isVisible">
        <div class="popup">
            <p>{{ message }}</p>
            <button @click="confirm">OK</button>
            <button @click="cancel">Cancel</button>
        </div>
    </div>
</template>

<script>
export default {
    props: {
        isVisible: Boolean,
        message: String,
    },
    methods: {
        confirm() {
            this.$emit("confirmed");
        },
        cancel() {
            this.$emit("cancelled");
        },
    },
};
</script>

<style scoped>
.popup-overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    background-color: rgba(0, 0, 0, 0.5);
}

.popup {
    background-color: white;
    padding: 20px;
    border-radius: 10px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
    animation: zoomOut 0.3s ease-out;
}

@keyframes zoomOut {
    from {
        transform: scale(0);
        opacity: 0;
    }

    to {
        transform: scale(1);
        opacity: 1;
    }
}
</style>
