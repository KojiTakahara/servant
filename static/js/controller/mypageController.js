'use strict';

var app = angular.module('mypageCtrl', []);

app.controller('editDeckController', ['$scope', '$location', function($scope, $location) {
  $scope.cardNums = [1, 2, 3, 4];
  $scope.deck = {
    lrig: [],
    main: []
  };

  $scope.addCard = function(index) {
    var card = $scope.cards[index],
        deck = $scope.deck;
    card.num = 1;
    if (isLrigDeck(card) && !isContain(deck.lrig, card)) {
      deck.lrig.push(card);
    } else if (isMainDeck(card) && !isContain(deck.main, card)) {
      deck.main.push(card);
    }
  };

  $scope.removeCard = function(category, index) {
    if (isLrigDeck({Category: category})) {
      $scope.deck.lrig.splice(index, 1);
    } else if (isMainDeck({Category: category})) {
      $scope.deck.main.splice(index, 1);
    }
  };

  var isContain = function(list, card) {
    var result = false;
    for (var i in list) {
      if (list[i].Id === card.Id) {
        result = true;
        break;
      }
    }
    return result;
  }

  var isLrigDeck = function(card) {
    return card.Category === 'ルリグ' || card.Category === 'アーツ';
  };
  var isMainDeck = function(card) {
    return card.Category === 'シグニ' || card.Category === 'スペル';
  };

  /** 最後のカードでtabが押されたらinputに戻すやつ **/
  $scope.returnCursor = function(bool, event) {
    if (bool && event.which === 9) {
      console.log(1);
      angular.element('.search-query').selected;
    }
  };

  $scope.scopes = [{
    id: 'PRIVATE',
    name: '非公開'
  }, {
    id: 'SELECT',
    name: '限定公開'
  }, {
    id: 'PUBLIC',
    name: '公開'
  }];


  $scope.cards = [
    {"Id":329,"KeyName":"","Burst":"","Bursted":false,"Category":"アーツ","Color":"blue","Constraint":"エルドラ限定","CostBlack":-1,"CostBlue":1,"CostColorless":1,"CostGreen":-1,"CostRed":-1,"CostWhite":-1,"Expansion":"WD06","Flavor":"ばっちん！","Guard":"","Illus":"パトリシア","Image":"/products/wixoss/images/card/WD06/WD06-006.jpg","Level":-1,"Limit":-1,"Name":"クロス・クラッシュ・フラッシュ","NameKana":"クロスクラッシュフラッシュ","No":6,"Power":-1,"Reality":"ST","SearchText":"","Text":"使用タイミング【メインフェイズ】,\n対戦相手はライフクロスの一番上を公開する。そのカードが【ライフバースト】を持たない場合、それをトラッシュに置く。","Type":"","Url":"http://www.takaratomy.co.jp/products/wixoss/card/card_detail.php?id=329","ParentKeyName":""},
    {"Id":360,"KeyName":"","Burst":"","Bursted":false,"Category":"ルリグ","Color":"white","Constraint":"","CostBlack":-1,"CostBlue":-1,"CostColorless":-1,"CostGreen":-1,"CostRed":-1,"CostWhite":0,"Expansion":"PR","Flavor":"いこっ！　～タマ～","Guard":"","Illus":"ＰＯＰ","Image":"/products/wixoss/images/card/PR/PR-001.jpg","Level":0,"Limit":0,"Name":"新月の巫女　タマヨリヒメ ","NameKana":"シンゲツノミコタマヨリヒメ","No":1,"Power":-1,"Reality":"PR","SearchText":"","Text":"","Type":"タマ","Url":"http://www.takaratomy.co.jp/products/wixoss/card/card_detail.php?id=360","ParentKeyName":""},
    {"Id":324,"KeyName":"","Burst":"","Bursted":false,"Category":"ルリグ","Color":"blue","Constraint":"","CostBlack":-1,"CostBlue":3,"CostColorless":-1,"CostGreen":-1,"CostRed":-1,"CostWhite":-1,"Expansion":"WD06","Flavor":"ジャッジャーン！エルドラゴールデンハンマー！　～エルドラ～","Guard":"","Illus":"ナダレ","Image":"/products/wixoss/images/card/WD06/WD06-001.jpg","Level":4,"Limit":11,"Name":"エルドラ＝マークⅣ´","NameKana":"エルドラマークフォーダッシュ","No":1,"Power":-1,"Reality":"ST","SearchText":"","Text":"[常]：あなたのライフバーストが発動するたび、カードを1枚引く。,\n[常]：あなたのライフクロスにカード1枚が加えられるたび、あなたは(青)(青)(青)を支払ってもよい。そうした場合、対戦相手のシグニ1体をバニッシュする。","Type":"エルドラ","Url":"http://www.takaratomy.co.jp/products/wixoss/card/card_detail.php?id=324","ParentKeyName":""},
    {"Id":351,"KeyName":"","Burst":"カードを1枚引く。","Bursted":true,"Category":"シグニ","Color":"white","Constraint":"イオナ限定","CostBlack":-1,"CostBlue":-1,"CostColorless":-1,"CostGreen":-1,"CostRed":-1,"CostWhite":-1,"Expansion":"WD07","Flavor":"人が作りし精械、迷宮の名を冠す。","Guard":"","Illus":"arihato","Image":"/products/wixoss/images/card/WD07/WD07-010.jpg","Level":2,"Limit":-1,"Name":"コードメイズ　バベル","NameKana":"コードメイズバベル","No":10,"Power":3000,"Reality":"ST","SearchText":"","Text":"[常]：対戦相手がシグニを配置する場合、可能ならばこのシグニの正面に配置しなければならない。,\n[起]手札を1枚捨てる(Ｔ)：あなたのデッキからレベル3以下の黒のシグニ1枚を探して公開し手札に加える。その後、デッキをシャッフルする。","Type":"精械：迷宮","Url":"http://www.takaratomy.co.jp/products/wixoss/card/card_detail.php?id=351","ParentKeyName":""},
    {"Id":355,"KeyName":"","Burst":"","Bursted":false,"Category":"シグニ","Color":"black","Constraint":"","CostBlack":-1,"CostBlue":-1,"CostColorless":-1,"CostGreen":-1,"CostRed":-1,"CostWhite":-1,"Expansion":"WD07","Flavor":"不可視の現実、謎の芸。","Guard":"","Illus":"よこえ","Image":"/products/wixoss/images/card/WD07/WD07-014.jpg","Level":2,"Limit":-1,"Name":"コードアンチ　モア","NameKana":"コードアンチモア","No":14,"Power":5000,"Reality":"ST","SearchText":"","Text":"[常]：あなたの場に他の＜古代兵器＞のシグニがあるかぎり、このシグニのパワーは8000になる。","Type":"精械：古代兵器","Url":"http://www.takaratomy.co.jp/products/wixoss/card/card_detail.php?id=355","ParentKeyName":""},
    {"Id":352,"KeyName":"","Burst":"カードを1枚引く。","Bursted":true,"Category":"シグニ","Color":"white","Constraint":"イオナ限定","CostBlack":-1,"CostBlue":-1,"CostColorless":-1,"CostGreen":-1,"CostRed":-1,"CostWhite":-1,"Expansion":"WD07","Flavor":"スットレート！　～凱旋～","Guard":"","Illus":"かざあな","Image":"/products/wixoss/images/card/WD07/WD07-011.jpg","Level":1,"Limit":-1,"Name":"コードメイズ　凱旋","NameKana":"コードメイズガイセン","No":11,"Power":1000,"Reality":"ST","SearchText":"","Text":"[常]：対戦相手がシグニを配置する場合、可能ならばこのシグニの正面に配置しなければならない。,\n[起]手札を1枚捨てる(Ｔ)：あなたのデッキからレベル2以下の黒のシグニ1枚を探して公開し手札に加える。その後、デッキをシャッフルする。","Type":"精械：迷宮","Url":"http://www.takaratomy.co.jp/products/wixoss/card/card_detail.php?id=352","ParentKeyName":""},
    {"Id":350,"KeyName":"","Burst":"あなたのデッキから黒のシグニ1枚を探して公開し手札に加える。その後、デッキをシャッフルする。","Bursted":true,"Category":"シグニ","Color":"white","Constraint":"イオナ限定","CostBlack":-1,"CostBlue":-1,"CostColorless":-1,"CostGreen":-1,"CostRed":-1,"CostWhite":-1,"Expansion":"WD07","Flavor":"やあ、あたらしい世界へようこそ！　～金字塔～","Guard":"","Illus":"しおぼい","Image":"/products/wixoss/images/card/WD07/WD07-009.jpg","Level":3,"Limit":-1,"Name":"コードメイズ　金字塔","NameKana":"コードメイズピラミッド","No":9,"Power":7000,"Reality":"ST","SearchText":"","Text":"[常]：対戦相手がシグニを配置する場合、可能ならばこのシグニの正面に配置しなければならない。,\n[起](Ｔ)：デッキの上からカードを3枚見る。その中から白または黒のシグニ1枚を公開して手札に加える。その後、残りのカードを好きな順番でデッキの一番上に戻す。","Type":"精械：迷宮","Url":"http://www.takaratomy.co.jp/products/wixoss/card/card_detail.php?id=350","ParentKeyName":""},
  ];
}]);

app.controller('mypageController', ['$scope', '$location', function($scope, $location) {
  $scope.decks = [];

  $scope.editDeck = function() {
    alert('編集処理を実装してね');
  };

  $scope.deleteDeck = function(index) {
    alert('削除処理を実装してね ' + index);
  };

  $scope.createDeck = function() {
    $location.path('/mypage/deck/0');
  };

  var init = function() {
    for (var i = 0; i < 10; i++) {
      $scope.decks.push({
        Title: 'テストデッキ' + i,
        Id: 1000 + i,
        White: true,
        Red: true,
        Blue: true,
        Green: true,
        Black: true,
        Scope: 'SELECT',
        UpdatedAt: new Date()
      });
    }
  };
  init();

}]);