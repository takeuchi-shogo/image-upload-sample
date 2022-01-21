<script>
	import axios from 'axios/dist/axios';


	let status = ''

	let dropped = []

	let files = []
	let file

	let reader

	$: {
		if (files != []) {
			preview(files[0])
		}
	}

	function preview(file) {
		if (!file) {
			console.log('ファイルがありません')
			return
		}
		reader = new FileReader()
		let image = new Image()

		reader.onloadend = () => {
			image.onload = () => {
				//　画像種別の取得
				let imageType = image.src.substring(5, image.src.indexOf(';'))
				// キャンバスの作成
				let canvas = document.getElementById('canvas')
				canvas.width = 280
				canvas.height = 200
				// リサイズ
				let ctx = canvas.getContext('2d')
				ctx.drawImage(image, 0, 0, 280, 200)
				// imgタグのsrcプロパティにURIを渡す
				// img.src = canvas.toDataURL(imageType, 1)
			}
			// イメージオブジェクトにURIを渡して画像として扱う　<img src="{ ここにURLが入り表示される }">
			image.src = reader.result
		}

		reader.readAsDataURL(file)
	}

	function handleDragEventer(e) {
		status = "You are dragging over the " + e
			.target
			.getAttribute('id')
	}

	function handleDragDrop(e) {
		e.preventDefault()
		var element_id = e
			.dataTransfer
			.getDate('text')
		dropped = dropped.concat(element_id)
		dropped_in = true
		status = "You Droped"
	}

	function log() {
		console.log('files: ', files)
		file = files.item(0)
		console.log('file: ', file)
		preview(files[0])
		let params = new FormData()
		params.append('file', files[0])
		let header = { "Content-Type": "multipart/form-data" }
		let url = '/images'
		console.log('params: ', ...params.entries())
		let axiosAPI = axios.create({
			baseURL: 'http://localhost:8080',
		})
		axiosAPI.post(url, params, { headers: header })
			.then((res) => {
				console.log('res: ', res)
			})
			.catch((e) => {
				console.log(e)
			})
	}
</script>

<main>
	<h2 id="app_status">Drag Status: { status }</h2>
	<h1>Drop zone</h1>

	<div
		on:dragenter="{handleDragEventer}"
		
		on:drop="{handleDragDrop}"
		id="drop_zone"
		ondragover="return false"
	>
		<div class="object">
			
		</div>
	</div>

	<!-- <form> -->
		<label class="load-input">
			ファイルを選択
			<input type="file" name="example" id="example" accept="image/png, image/jpeg" bind:files>
		</label>
		<button on:click="{ log }">送信</button>
	<!-- </form> -->

	<div class="image-preview" id="image-preview">
		<!-- <img src="" id="img" alt=""/> -->
		<canvas id="canvas"></canvas>
	</div>
</main>

<style>
	#drop_zone{
		background-color: #eee;
		border: #999 1px solid ;
		width: 280px;
		height: 200px;
		padding: 8px;
		font-size: 19px;
	}

	.load-input {
		display: inline-block;
		cursor: pointer; /* カーソルを指に */
		margin: 1em 0; /* まわりの余白 */
		padding: .7em 1em; /* 文字まわりの余白 */
		line-height: 1.4; /* 行間 */
		background: #3e8bff; /* 背景色 */
		color: #FFF; /* 文字色 */
		font-size: 0.95em; /* フォントサイズ */
		border-radius: 2.5em; /* 角の丸み */
		transition: 0.2s; /* ホバーをなめらかに */
	}

	.load-input input {
		display: none;
	}

	.image-preview {
		height: 200px;
		width: 280px;
		border: #999 1px solid;
	}

</style>