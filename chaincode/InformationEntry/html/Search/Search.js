/*
 * @Author: your name
 * @Date: 2021-12-14 10:06:08
 * @LastEditTime: 2021-12-14 10:36:58
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \blockchain_accounting2020_1.2d:\fabric\fabric-sh\chaincode\InformationEntry\html\js\index.js
 */
$(function() {
    console.log('页面加载完成');
    // 录入学生信息保存
    $("#saveBtn").click(function(){
        var searchInput = $('#searchInput').val();
        console.log('学生姓名：' + searchInput);

        var params = {
            'searchInput': searchInput,
        }

        console.log(params);

    });
});