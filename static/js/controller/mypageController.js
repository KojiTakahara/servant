'use strict';

var app = angular.module('mypageCtrl', [
  'apiService',
  'deckService',
]);

app.controller('editDeckController', ['$scope', '$stateParams', '$location', '$filter', '$anchorScroll', 'cardService', 'deckService', function($scope, $stateParams, $location, $filter, $anchorScroll, cardService, deckService) {
  $scope.alerts = [];
  $scope.cardNums = [1, 2, 3, 4];
  $scope.deck = {
    Lrig: [],
    Main: [],
    Scope: 'PRIVATE',
  };
  $scope.searchText = '';

  var init = function() {
    var id = $stateParams.id;
    if (id === '0') {
      return;
    }
    cardService.getDeckById(id).then(function(res) {
      $scope.deck = res;
    }, function(err) {
      alert('error');
    });
  };
  init();

  $scope.addCard = function(card) {
    var deck = $scope.deck;
    if (isLrigDeck(card) && !isContain(deck.Lrig, card)) {
      card.Num = 1;
      deck.Lrig.push(card);
    } else if (isMainDeck(card) && !isContain(deck.Main, card)) {
      card.Num = 1;
      deck.Main.push(card);
    }
  };

  $scope.removeCard = function(category, index) {
    if (isLrigDeck({Category: category})) {
      $scope.deck.Lrig.splice(index, 1);
    } else if (isMainDeck({Category: category})) {
      $scope.deck.Main.splice(index, 1);
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

  $scope.countLrig = function() {
    return deckService.countLrig($scope.deck);
  };

  $scope.countMain = function() {
    return deckService.countMain($scope.deck);
  };

  $scope.countLifeBurst = function() {
    return deckService.countBurst($scope.deck);
  };

  /** デッキ保存処理 **/
  $scope.save = function() {
    $scope.alerts = deckService.validate($scope.deck);
    if (0 < $scope.alerts.length) {
      return;
    }
    $scope.deck.Main = $filter('orderBy')($scope.deck.Main, $scope.deck.mainPredicate, $scope.deck.mainReverse);
    $scope.deck.Lrig = $filter('orderBy')($scope.deck.Lrig, $scope.deck.lrigPredicate, $scope.deck.lrigReverse);
    $anchorScroll.yOffset = 0;
    cardService.saveDeck($scope.deck).then(function(res) {
      $scope.alerts.push({ type: 'success', msg: '保存しました。' });
      $anchorScroll();
    }, function(err) {
      $scope.alerts.push({ type: 'danger', msg: '保存できませんでした。' });
      $anchorScroll();
    });
  };

  $scope.sort = function(list, predicate) {
    // if (list === 'lrig') {
    //   $scope.deck.lrigPredicate = predicate;
    //   $scope.deck.lrigReverse = !$scope.deck.lrigReverse;
    // }
    // if (list === 'main') {
    //   $scope.deck.mainPredicate = predicate;
    //   $scope.deck.mainReverse = !$scope.deck.mainReverse;
    // }
  };

  $scope.closeAlert = function(index) {
    $scope.alerts.splice(index, 1);
  };

  $scope.scopes = [
    'PRIVATE',
    // 'SELECT',
    'PUBLIC'
  ];

  $scope.deselect = function(text) {
    $scope.searchText = text;
  };

  $scope.selectTab = function(category) {
    if (category === 'lrig') {
      if ($scope.lrigList === undefined) {
        cardService.search({ category: 'ルリグ' }).then(function(res) { $scope.lrigList = res; });
      }
      $scope.lrigSearchText = $scope.searchText;
    }
    if (category === 'arts') {
      if ($scope.artsList === undefined) {
        cardService.search({ category: 'アーツ' }).then(function(res) { $scope.artsList = res; });
      }
      $scope.artsSearchText = $scope.searchText;
    }
    if (category === 'signi') {
      if ($scope.signiList === undefined) {
        cardService.search({ category: 'シグニ' }).then(function(res) { $scope.signiList = res; });
      }
      $scope.signiSearchText = $scope.searchText;
    }
    if (category === 'spell') {
      if ($scope.spellList === undefined) {
        cardService.search({ category: 'スペル' }).then(function(res) { $scope.spellList = res; });
      }
      $scope.spellSearchText = $scope.searchText;
    }
  };
  $scope.selectTab('lrig');

}]);

app.controller('mypageController', ['$scope', '$location', '$window', 'cardService', function($scope, $location, $window, cardService) {
  $scope.decks = [];

  $scope.editDeck = function() {
    alert('編集処理を実装してね');
  };

  $scope.deleteDeck = function(index) {
    var deck = $scope.decks[index];
    if ($window.confirm(deck.Title + 'を削除しますか？')) {
      cardService.deleteDeck(deck).then(function(res) {
        alert('削除しました');
      }, function(err) {
        alert('削除に失敗しました');
      });
    }
  };

  $scope.createDeck = function() {
    $location.path('/mypage/deck/0');
  };

  var init = function() {
    cardService.getDeckByUserId("dm_plateau").then(function(res) {
      $scope.decks = res;
    }, function(err) {
      //
    });
  };
  init();

}]);