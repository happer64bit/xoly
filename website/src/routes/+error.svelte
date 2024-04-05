<script>
    import { page } from '$app/stores'
</script>

<div>
    <div class="h-[80dvh] flex items-center justify-center">
        <div>
            <h1 class="text-2xl">Error</h1>
            <p>
                Caught Error: <code>{$page.error?.message}</code>
            </p>
        </div>
    </div>
</div>