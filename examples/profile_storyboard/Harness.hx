import domain.StoryCard;
import haxe.ds.List;
import profile.StoryboardRuntime;

class Harness {
  static function makeTags(a:String, ?b:String):List<String> {
    var tags = new List<String>();
    tags.add(a);
    if (b != null) {
      tags.add(b);
    }
    return tags;
  }

  static function joinStringList(values:List<String>, separator:String):String {
    var out = "";
    var first = true;
    var count = values.length;
    var i = 0;
    while (i < count) {
      var raw = values.pop();
      if (raw == null) {
        break;
      }
      var value:String = cast raw;
      if (!first) {
        out += separator;
      }
      out += value;
      values.add(value);
      first = false;
      i++;
    }
    return out;
  }

  static function card(id:Int, title:String, points:Int, tags:List<String>):StoryCard {
    return new StoryCard(id, title, points, tags);
  }

  public static function buildCards():List<StoryCard> {
    var cards = new List<StoryCard>();
    cards.add(card(1, "Ship profile docs", 3, makeTags("docs", "profiles")));
    cards.add(card(2, "Backfill regression snapshots", 5, makeTags("tests")));
    cards.add(card(3, "Wire release artifacts", 5, makeTags("ci", "release")));
    return cards;
  }

  static function baselineSummary(cards:List<StoryCard>):String {
    var totalPoints = 0;
    var open = 0;
    var count = cards.length;
    var i = 0;
    while (i < count) {
      var value = cards.pop();
      if (value == null) {
        break;
      }
      var card:StoryCard = cast value;
      totalPoints += card.points;
      open++;
      cards.add(card);
      i++;
    }
    return "cards=" + open + ",points=" + totalPoints + ",open=" + open;
  }

  static function formatCards(cards:List<StoryCard>, runtime:StoryboardRuntime):String {
    var out = "";
    var firstCard = true;
    var cardCount = cards.length;
    var i = 0;
    while (i < cardCount) {
      var cardValue = cards.pop();
      if (cardValue == null) {
        break;
      }
      var card:StoryCard = cast cardValue;

      var tags = new List<String>();
      var tagCount = card.tags.length;
      var j = 0;
      while (j < tagCount) {
        var tagValue = card.tags.pop();
        if (tagValue == null) {
          break;
        }
        var tag:String = cast tagValue;
        tags.add(runtime.highlightTag(tag));
        card.tags.add(tag);
        j++;
      }

      if (!firstCard) {
        out += ";";
      }
      out += "#" + card.id + ":" + runtime.decorateTitle(card.title) + ":p" + card.points + ":" + joinStringList(tags, "|");
      firstCard = false;

      cards.add(card);
      i++;
    }
    return out;
  }

  public static function render(runtime:StoryboardRuntime):String {
    var cards = buildCards();
    var velocity = "off";
    if (runtime.supportsVelocityHint()) {
      velocity = "on";
    }

    return "profile=" + runtime.profileId()
      + "\nbaseline=" + baselineSummary(cards)
      + "\nview=" + formatCards(cards, runtime)
      + "\nextra=" + runtime.extraSignal(cards)
      + "\nvelocity_hint=" + velocity;
  }

  public static function assertContract(runtime:StoryboardRuntime):String {
    var cards = buildCards();
    var baseline = baselineSummary(cards);
    if (baseline != "cards=3,points=13,open=3") {
      throw "baseline drift: " + baseline;
    }
    var extra = runtime.extraSignal(cards);
    if (extra == null || extra == "") {
      throw "missing extra signal";
    }
    return "OK " + runtime.profileId();
  }
}
