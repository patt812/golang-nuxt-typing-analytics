<script setup lang="ts">
const { apiBase } = useRuntimeConfig().public
const results = ref<string[][] | null>()

const kana = ref('')
const analyze = async () => {
  const res = await useFetch<AnalyzeResponse | null>(`${apiBase}/analyze?kana=${kana.value}`)
  results.value = res.data.value?.patterns
}
</script>

<template>
  <form @submit.prevent="analyze">
    <input v-model="kana" type="text" name="kana" id="kana" pattern="[\u3040-\u309Fヴヵヶゃゅょゎゐゑ]+" title="ひらがなのみを入力してください" />
    <button type="submit" :disabled="!kana.length">Submit</button>
  </form>

  <div v-if="results?.length">
    <div>{{ results.length }} patterns</div>

    <div v-for="(result, index) in results" :key="index">
      <div>
        <span>{{ index + 1 }}:&nbsp;</span>
        <span>{{ result }}</span>
      </div>
    </div>
  </div>
</template>
