<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Description</title>
    <script type="text/javascript" src="/{{.baseurl}}/js/jquery.js"></script>
    <script type="text/javascript" src="/{{.baseurl}}/js/effects-20090707.js"></script>
    <script type="text/javascript" src="/{{.baseurl}}/js/header.js"></script>
    <script type="text/javascript" src="/{{.baseurl}}/js/messager.js"></script>
    <script type="text/javascript" src="/{{.baseurl}}/js/jquery.boxy.js"></script>
    <script type="text/javascript" src="/{{.baseurl}}/js/jquery-ui-1.8.21.custom.min.js"></script>

    <!--<link type="text/css" rel="stylesheet" href="/{{.baseurl}}/css/common.css"/>-->
    <!--<link type="text/css" rel="stylesheet" href="/{{.baseurl}}/css/boxy.css"/>-->
    <!--<link type="text/css" rel="stylesheet" href="/{{.baseurl}}/css/jquery_ui/jquery-ui-1.8.21.custom.css"/>-->
    <!--<link type="text/css" rel="stylesheet" href="/{{.baseurl}}/css/home2010.css"/>-->
    <!--&lt;!&ndash;FontAwesome&ndash;&gt;-->
    <!--<link type="text/css" rel="stylesheet" href="/{{.baseurl}}/plugin/font-awesome/font-awesome.min.css"/>-->

    <style>
        img
        {
            width: 100%;
        }
    </style>


    <script>
    function checkKC()
    {
    if ($("#viewkc").get(0).checked==true){$(".p_kc").show();$("#viewbyKC").addClass("viewindex_curr");createCookie("viewkc",1);}
    else{$(".p_kc").hide();$("#viewbyKC").removeClass("viewindex_curr");createCookie("viewkc",0);}
    }
    $(function(){
    showSort("SortList");
    /*ViewMode*/
    $(".viewindex").click(function(){
    $(this).siblings().each(function(){$(this).removeClass("viewindex_curr");})
    if (this.className=="viewindex"&&this.id!="viewbyKC"){$(this).toggleClass("viewindex_curr");}else{return;}
    });
    $(".viewindex dl").hover(function(){$(this).addClass('hover');},function(){$(this).removeClass('hover');});

    $(".viewindex dd").hover(function(){$(this).css({background:"#FFF3BF"});},function(){$(this).css({background:"#ffffff"});}).click(function(){$(this).parent().get(0).className="";if ($(this).parent().parent().attr("id")=="viewbyKC"){createCookie("viewkc",1);}
    });

    $("#viewkc").click(function(){checkKC();})
    if (readCookie("viewkc")==1){$("#viewkc").get(0).checked=true;checkKC();}
    var url = window.location.href ;
    var param = url.split("-")[11];
    if(param=="4" && url.split("-")[2]!="0")
    {
    $(".p_Review span").each(function(){
    if ($(this).attr("rel")!=0)
    {
    $(this).show();
    }
    });
    }
    });

    function search(id){
    search_submit();
    }
    function login(){
    return false;
    }
    function search_submit(){
    if($('.S_input').find("input").attr("value")==''){alert("请输入查询关键词！");return false;}
    if($('.S_input').find("input").css("color")=="rgb(153, 153, 153)"){return false;}
    document.searchForm.submit();

    }

    //Nav背景切换
    var NavIndex = 0;
    $(function(){
    $("#Nav li:nth-child("+NavIndex+ ")").siblings().removeClass("curr").end().addClass("curr");
    });

    //建立全局变量，取得site_url
    global_siteUrl = '/{{.baseurl}}/';

    </script>
    <link type="text/css" rel="stylesheet" href="/{{.baseurl}}/css/aixin_dongyueweb.css"/>

            <script type="text/javascript">
    var _gaq = _gaq || [];
    _gaq.push(['_setAccount', 'UA-29996992-1']);
    _gaq.push(['_trackPageview']);

    (function() {
        var ga = document.createElement('script'); ga.type = 'text/javascript'; ga.async = true;
        ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
        var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(ga, s);
    })();

    </script>


</head>
<!--<body style="background: #F4F6F8 url(/{{.baseurl}}/images/background/01.jpg) no-repeat top left;background-size:cover;filter: progid:DXImageTransform.Microsoft.AlphaImageLoader( src='/{{.baseurl}}/images/background/01.jpg', sizingMethod='scale')\9;">-->
<body>

<script type="text/javascript" src="/{{.baseurl}}/js/jquery.jcarousel.min.js"></script>
<script type="text/javascript" src="/{{.baseurl}}/js/jquery.tooltip.js"></script>
<link type="text/css" rel="stylesheet" href="/{{.baseurl}}/css/jquery.tooltip.css"/>
<script type="text/javascript">
    $(document).ready(function() {
        $('.level_3_tooltip').tooltip({
            track: true,
            delay: 0,
            fade: 250 ,
            extraClass: "pretty fancy"
        });
        $('img').width('100%');
        $('img').height('100%');
//        alert("change");
    });
</script>

<script type="text/javascript">
    function mycarousel_initCallback(carousel){
        $("#mycarousel img")[0].className="curr";
        $("#mycarousel li").mouseover(function(){
            $("#Product_BigImage img")[0].src=this.getElementsByTagName("img")[0].name;
            $("#Product_BigImage img")[0].jqimg=this.getElementsByTagName("img")[0].name;
            $(this).siblings().each(function(){
                this.getElementsByTagName("img")[0].className="";
            })
            this.getElementsByTagName("img")[0].className="curr";
        })
    };

    $(function(){
        if ($.browser.msie){
            mycarousel_initCallback();
        }else{
            jQuery("#mycarousel").jcarousel({initCallback:mycarousel_initCallback});
        };

        $("#rank-tab-1").jdTab();
        $("#rank-con-1 li").jdSonny();
        $("#rank-con-2 li").jdSonny();
        $("#rank-con-3 li").jdSonny();
    });
</script>

<div class="Main">
    <div class="right">
        <div class="Product_Intro">
            <div class="Product_Name">
                <h1 style='margin:10px 4px;'>{{.product.Name}}</h1>
                <div style='color:darkGreen;margin-left:30px;' >{{.product.Short_desc}}</div>
            </div>

            <div class="sbox_3" id="EFF_PINFO_Con_0">
                <ul class="Detail1">
                    <li class="w100">物品名称：{{.product.Name}}</li>
                    <li>上架时间：{{.product.On_sale_at}}</li>
                </ul>
                <div class="ProInfo">{{.desp}}</div>
            </div>

            <div id="suit"></div>

        </div>
    </div>
</div>





</body>
</html>