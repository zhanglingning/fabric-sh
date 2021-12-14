/*
 * @Author: your name
 * @Date: 2021-12-14 10:06:08
 * @LastEditTime: 2021-12-14 10:27:01
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \blockchain_accounting2020_1.2d:\fabric\fabric-sh\chaincode\InformationEntry\html\js\index.js
 */
$(function() {
    console.log('页面加载完成');
    // 录入学生信息保存
    $("#saveBtn").click(function(){
        var userName = $('#userName').val();
        console.log('学生姓名：' + userName);

        var userPhone = $('#userPhone').val();
        console.log('电话：' + userPhone);

        var sex = $('#sex').val();
        console.log('性别：' + sex);

        var userIdCard = $('#userIdCard').val();
        console.log('身份证号码：' + userIdCard);

        var userAddress = $('#userAddress').val();
        console.log('籍贯：' + userAddress);

        var userNation = $('#userNation').val();
        console.log('民族：' + userNation);

        var userSchool = $('#userSchool').val();
        console.log('学校：' + userSchool);


        var userEnterTime = $('#userEnterTime').val();
        console.log('入学日期：' + userEnterTime);

        var userMajor = $('#userMajor').val();
        console.log('专业：' + userMajor);

        var userFaculty = $('#userFaculty').val();
        console.log('院系：' + userFaculty);

        var params = {
            'userName': userName,
            'userPhone': userPhone,
            'sex': sex,
            'userIdCard': userIdCard,
            'userAddress': userAddress,
            'userNation': userNation,
            'userSchool': userSchool,
            'userEnterTime': userEnterTime,
            'userMajor': userMajor,
            'userFaculty': userFaculty,
        }

        console.log(params);

    });
});