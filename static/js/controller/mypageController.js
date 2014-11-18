'use strict';

var app = angular.module('mypageCtrl', [
  'apiService',
]);

app.controller('editDeckController', ['$scope', '$location', 'cardService', function($scope, $location, cardService) {
  $scope.alerts = [];
  $scope.cardNums = [1, 2, 3, 4];
  $scope.deck = {
    lrig: [],
    main: [],
    Scope: 'PRIVATE',
  };

  $scope.addCard = function(card) {
    var deck = $scope.deck;
    if (isLrigDeck(card) && !isContain(deck.lrig, card)) {
      card.num = 1;
      deck.lrig.push(card);
    } else if (isMainDeck(card) && !isContain(deck.main, card)) {
      card.num = 1;
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
      angular.element('#cursor').focus();
    }
  };

  $scope.isEmpty = function(obj) {
    return obj === undefined || obj === "";
  };

  /** デッキ保存処理 **/
  $scope.save = function() {
    var lrigDeck = $scope.deck.lrig,
        mainDeck = $scope.deck.main,
        lrigNum = 0,
        mainNum = 0;
    $scope.alerts = [];
    for (var i in lrigDeck) {
      lrigNum += lrigDeck[i].num;
    }
    for (var i in mainDeck) {
      mainNum += mainDeck[i].num;
    }
    if ($scope.isEmpty($scope.deck.Title)) {
      $scope.alerts.push({ type: 'warning', msg: 'デッキ名を入力してください。' });
    }
    if (10 < lrigNum) {
      $scope.alerts.push({ type: 'warning', msg: 'ルリグデッキは合計10枚までにしてください。' });
    }
    if (mainNum !== 40) {
      $scope.alerts.push({ type: 'warning', msg: 'メインデッキは合計40枚にしてください。' });
    }
  };

  $scope.sort = function(list, predicate) {
    if (list === 'lrig') {
      $scope.deck.lrigPredicate = predicate;
      $scope.deck.lrigReverse = !$scope.deck.lrigReverse;
    }
    if (list === 'main') {
      $scope.deck.mainPredicate = predicate;
      $scope.deck.mainReverse = !$scope.deck.mainReverse;
    }
  };

  $scope.closeAlert = function(index) {
    $scope.alerts.splice(index, 1);
  };

  $scope.scopes = {
    'PRIVATE': '非公開',
    'SELECT': '限定公開',
    'PUBLIC': '公開',
  };

  cardService.search({ category: 'ルリグ' }).then(function(res) { $scope.lrigList = res; });
  cardService.search({ category: 'アーツ' }).then(function(res) { $scope.artsList = res; });
  cardService.search({ category: 'シグニ' }).then(function(res) { $scope.signiList = res; });
  cardService.search({ category: 'スペル' }).then(function(res) { $scope.spellList = res; });

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