let ws_protocol = document.location.protocol == "https:" ? "wss" : "ws"

const websocketHeartbeatJsOptions = {
	url: ws_protocol + "://"+ window.location.host +"/ws",
	pingTimeout: 15000,
	pongTimeout: 10000,
	reconnectTimeout: 2000,
	pingMsg: "heartbeat"
}

let websocketHeartbeatJs = new WebsocketHeartbeatJs(websocketHeartbeatJsOptions);

let ws = websocketHeartbeatJs;

function _time(time = +new Date()) {
	var date = new Date(time + 8 * 3600 * 1000); // 增加8小时
	return date.toJSON().substr(0, 19).replace('T', ' ');
}

function WebSocketConnect(userInfo,toUserInfo = null) {
	if ("WebSocket" in window) {

		if ( userInfo.uid <= 0 )
		{
			console.log('参数错误，请刷新页面重试');return false;
		}
		let send_data = JSON.stringify({
			"status": toUserInfo ? 5 : 1,
			"data": {
				"uid": userInfo.uid,
				"room_id": userInfo.room_id,
				"avatar_id": userInfo.avatar_id,
				"username": userInfo.username,
				"to_user": toUserInfo
			}
		})

		ws.onopen = function () {
			ws.send(send_data);
		};

		if ( toUserInfo )
		{
			let to_user_send_data = JSON.stringify({
				"status": toUserInfo ? 5 : 1,
				"data": {
					"uid": toUserInfo.uid,
					"room_id": toUserInfo.room_id,
					"avatar_id": toUserInfo.avatar_id,
					"username": toUserInfo.username,
					"to_user": toUserInfo
				}
			})
			ws.onopen = function () {
				ws.send(to_user_send_data);
				console.log("to_user_send_data 发送数据", to_user_send_data)
			};
		}


		let chat_info = $('.main .chat_info')

		ws.onmessage = function (evt) {
			var received_msg = JSON.parse(evt.data);

			// let myDate = new Date();
			// let time = myDate.toLocaleDateString() + myDate.toLocaleTimeString()
			let time = _time(received_msg.data.time)

			switch(received_msg.status)
			{
				case 1:
					chat_info.html(chat_info.html() +
						'<li class="systeminfo"> <span>' +
						"【" +
						received_msg.data.username +
						"】" +
						time +
						" 加入了房间" +
						'</span></li>');
					break;
				case 2:
					chat_info.html(chat_info.html() +
						'<li class="systeminfo"> <span>' +
						"【" +
						received_msg.data.username +
						"】" +
						time +
						" 离开了房间" +
						'</span></li>');
					break;
				case 3:
					if ( received_msg.data.uid != userInfo.uid && !isPrivateChat())
					{
						chat_info.html(chat_info.html() +
							'<li class="left"><img src="/static/images/user/' +
							received_msg.data.avatar_id +
							'.png" alt=""><b>' +
							received_msg.data.username +
							'</b><i>' +
							time +
							'</i><div class="aaa">' +
							received_msg.data.content +
							'</div></li>');
					}
					break;
				case -1:
					ws.close() // 主动close掉
					console.log("client 连接已关闭...");
					break;
				case 4:
					$('.popover-title').html('在线用户 '+ received_msg.data.count +' 人')

					$.each(received_msg.data.list,function (index, value) {

						if ( received_msg.data.uid == value.uid )
						{
							// 禁止点击
							$('.ul-user-list').html($('.ul-user-list').html() +
								'<li  style="pointer-events: none;" class="li-user-item" data-uid='+ value.uid +' data-username='+ value.username +' data-room_id='+ value.room_id +' data-avatar_id='+ value.avatar_id +'  ><img src="/static/images/user/' +
								value.avatar_id +
								'.png" alt=""><b>' + " " +
								value.username +
								'</b>' +
								'</li>'
							)
						}else{
							$('.ul-user-list').html($('.ul-user-list').html() +
								'<li  class="li-user-item" data-uid='+ value.uid +' data-username='+ value.username +' data-room_id='+ value.room_id +' data-avatar_id='+ value.avatar_id +'  ><img src="/static/images/user/' +
								value.avatar_id +
								'.png" alt=""><b>' + " " +
								value.username +
								'</b>' +
								'</li>'
							)
						}

					})
					break;
				case 5:
					// 私聊通知
					if (!isPrivateChat())
					{
						layer.msg(received_msg.data.username+'：'+ received_msg.data.content);
					}
					break;
				default:
			}
		};

		ws.onclose = function () {
			chat_info.html(chat_info.html() +
				'<li class="systeminfo"> <span>' +
				"与服务器连接断开，请刷新页面重试" +
				'</span></li>');
			console.log("serve 连接已关闭... " + _time());
		};
	} else {
		// 浏览器不支持 WebSocket
		console.log("您的浏览器不支持 WebSocket!");
	}
}

$(document).ready(function() {

	$('#userinfo_sub').click(function (event) {
		var userName = $('.rooms .user_name input').val(); // 用户昵称
		var userPortrait = $('.rooms .user_portrait img').attr('portrait_id'); // 用户头像id
		if (userName == '') { // 如果不填用户昵称，就是以前的昵称
			userName = $('.rooms .user_name input').attr('placeholder');
		}


		$('.userinfo a b').text(userName); // 修改标题栏的用户昵称
		$('.rooms .user_name input').val(''); // 昵称输入框清空
		$('.rooms .user_name input').attr('placeholder', userName); // 昵称输入框默认显示用户昵称
		$('.topnavlist .popover').not($(this).next('.popover')).removeClass('show'); // 关掉用户面板
		$('.clapboard').addClass('hidden'); // 关掉模糊背景
	});

	$('.theme img').click(function (event) {
		var theme_id = $(this).attr('theme_id');
		$('.clapboard').click(); // 关掉用户模糊背景

		$('body').css('background-image', 'url(images/theme/' + theme_id + '_bg.jpg)'); // 设置背景
	});

	$(document).on('click', '.a-user-list', function (e) {
		$('.ul-user-list').html('')
		let send_data = JSON.stringify({
			"status": 4,
			"data": {
				"uid": parseInt($('.room').attr('data-uid')),
				"username": $('.room').attr('data-username'),
				"avatar_id": $('.room').attr('data-avatar_id'),
				"room_id": $('.room').attr('data-room_id'),
			}
		})
		ws.send(send_data);
	})

	$('.imgFileBtn').change(function (event) {

		var formData = new FormData();
		formData.append('file', $(this)[0].files[0]);
		$.ajax({
			url: '/img-kr-upload',
			type: 'POST',
			cache: false,
			data: formData,
			processData: false,
			contentType: false
		}).done(function (res) {
			console.log(res)

			var str = '<img src="' + res.data.url + '" />'

			let to_uid = "0"
			let status = 3
			if (isPrivateChat()) {
				// 私聊
				to_uid = getQueryVariable("uid")
				status = 5
			}

			sends_message($('.room').attr('data-username'), $('.room').attr('data-avatar_id'), str); // sends_message(昵称,头像id,聊天内容);

			let send_data = JSON.stringify({
				"status": status,
				"data": {
					"uid": parseInt($('.room').attr('data-uid')),
					"username": $('.room').attr('data-username'),
					"avatar_id": $('.room').attr('data-avatar_id'),
					"room_id": $('.room').attr('data-room_id'),
					"image_url": res.data.url,
					"content": str,
					"to_uid": to_uid,
				}
			})

			console.log("send_data", send_data)
			ws.send(send_data);


			// 滚动条滚到最下面
			$('.scrollbar-macosx.scroll-content.scroll-scrolly_visible').animate({
				scrollTop: $('.scrollbar-macosx.scroll-content.scroll-scrolly_visible').prop('scrollHeight')
			}, 500);

			// 解决input上传文件选择同一文件change事件不生效
			event.target.value = ''
		}).fail(function (res) {
		});


	});

	$("#emojionearea2")[0].emojioneArea.setFocus()

	$('#subxx').click(function (event) {
		//var str = $('.text input').val(); // 获取聊天内容
		var str = $("#emojionearea2")[0].emojioneArea.getText() // 获取聊天内容
		str = str.replace(/\</g, '&lt;');
		str = str.replace(/\>/g, '&gt;');
		str = str.replace(/\n/g, '<br/>');
		str = str.replace(/\[em_([0-9]*)\]/g, '<img src="images/face/$1.gif" alt="" />');
		if (str != '') {

			let to_uid = "0"
			let status = 3
			if (isPrivateChat()) {
				// 私聊
				to_uid = getQueryVariable("uid")
				status = 5
			}


			sends_message($('.room').attr('data-username'), $('.room').attr('data-avatar_id'), str); // sends_message(昵称,头像id,聊天内容);

			let send_data = JSON.stringify({
				"status": status,
				"data": {
					"uid": parseInt($('.room').attr('data-uid')),
					"username": $('.room').attr('data-username'),
					"avatar_id": $('.room').attr('data-avatar_id'),
					"room_id": $('.room').attr('data-room_id'),
					"content": str,
					"to_uid": to_uid,
				}
			})

			ws.send(send_data);

			// 滚动条滚到最下面
			$('.scrollbar-macosx.scroll-content.scroll-scrolly_visible').animate({
				scrollTop: $('.scrollbar-macosx.scroll-content.scroll-scrolly_visible').prop('scrollHeight')
			}, 500);

		}

		$("#emojionearea2")[0].emojioneArea.setText("")
		$("#emojionearea2")[0].emojioneArea.setFocus()
	});
});
