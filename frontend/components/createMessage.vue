<template>
    <div>
        <h2 class="text-2xl">新規メッセージの送信</h2>
        <form class="my-5">
            <div class="my-1">
                <label for="space-id">送信先のスペース</label>
                <select data-te-select-init id="space-id" class="w-full h-[2em] border-2 rounded">
                    <option value="1">One</option>
                    <option value="2">Two</option>
                    <option value="3">Three</option>
                    <option value="4">Four</option>
                    <option value="5">Five</option>
                    <option value="6">Six</option>
                    <option value="7">Seven</option>
                    <option value="8">Eight</option>
                    <option value="9">Nine</option>
                    <option value="10">Ten</option>
                </select>
            </div>
            <div class="my-1">
                <label for="content">内容</label>
                <textarea id="content" class="w-full h-[10em] border-2 rounded"></textarea>
            </div>
        </form>
        <button class="rounded-lg bg-cyan-600 px-10 py-2 hover:bg-cyan-700">
            <p class="text-white">送信</p>
        </button>
    </div>
</template>

<script>
</script>

<style scoped>

</style>