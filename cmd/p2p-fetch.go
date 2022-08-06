package main

import (
	"context"
	"fmt"
	"github.com/slvic/p2p-fetch/internal/app"
	"os"
	"os/signal"
	"syscall"
)

func run(ctx context.Context) error {
	newApp, err := app.Initialize(ctx)
	if err != nil {
		return err
	}
	err = newApp.Run(ctx)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGKILL, syscall.SIGINT)
	defer cancel()
	if err := run(ctx); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "app run: %s\n", err.Error())
	}
}

//\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
//\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
//\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
//package main
//
//import (
//	"bytes"
//	"errors"
//	"fmt"
//	"io"
//	"strings"
//
//	"golang.org/x/net/html"
//)
//
//func GetNodeByTag(doc *html.Node, tagName string) (*html.Node, error) {
//	var htmlNode *html.Node
//	var crawler func(*html.Node)
//	crawler = func(node *html.Node) {
//		if node.Type == html.ElementNode && node.Data == tagName {
//			htmlNode = node
//			return
//		}
//		for child := node.FirstChild; child != nil; child = child.NextSibling {
//			crawler(child)
//		}
//	}
//	crawler(doc)
//	if htmlNode != nil {
//		return htmlNode, nil
//	}
//	return nil, errors.New("could not find a node by tag")
//}
//
//func GetNodeByAttrKey(doc *html.Node, attrKey, attrVal string) (*html.Node, error) {
//	var htmlNode *html.Node
//	var crawler func(*html.Node)
//	crawler = func(node *html.Node) {
//		if node.Type == html.ElementNode {
//			for _, attr := range node.Attr {
//				if attr.Key == attrKey && attr.Val == attrVal {
//					htmlNode = node
//					return
//				}
//			}
//		}
//		for child := node.FirstChild; child != nil; child = child.NextSibling {
//			crawler(child)
//		}
//	}
//	crawler(doc)
//	if htmlNode != nil {
//		return htmlNode, nil
//	}
//	return nil, errors.New("could not find a node by attribute key")
//}
//
//func renderNode(n *html.Node) string {
//	var buf bytes.Buffer
//	w := io.Writer(&buf)
//	html.Render(w, n)
//	return buf.String()
//}
//
//func getTableRowNodes(tbody *html.Node) ([]*html.Node, error) {
//	var htmlNodes []*html.Node
//
//	var crawler func(*html.Node)
//	crawler = func(node *html.Node) {
//		if node.Type == html.ElementNode && node.Data == "tr" {
//			htmlNodes = append(htmlNodes, node)
//		}
//		if node.NextSibling == nil {
//			return
//		}
//		crawler(node.NextSibling)
//	}
//	crawler(tbody.FirstChild)
//	if len(htmlNodes) != 0 {
//		return htmlNodes, nil
//	}
//	return nil, errors.New("could not find table rows")
//}
//
//func main() {
//	doc, _ := html.Parse(strings.NewReader(htm))
//	contentTable, err := GetNodeByAttrKey(doc, "id", "content_table")
//	node, err := GetNodeByTag(contentTable, "tbody")
//	tableRowNodes, err := getTableRowNodes(node)
//	if err != nil {
//		return
//	}
//
//	for _, tableRowNode := range tableRowNodes {
//		key, err := GetNodeByAttrKey(tableRowNode, `class`, `bj`)
//		if err != nil && key != nil {
//			return
//		}
//	}
//
//	if tableRowNodes != nil {
//
//	}
//
//	body := renderNode(node)
//	fmt.Println(body)
//}
//
//const htm = `<!DOCTYPE html>
//<html>
//<head>
//    <title></title>
//</head>
//
//      <table id="content_table">
//    <thead>
//      <tr>
//<td class="info"><img src="https://www.bestchange.com/images/update.png" id="update_image" class="" alt="Exchange rates" title="Updating the rates..." width="16" height="16"></td><td class="bj changer"><a href="https://www.bestchange.com/index.php?sort=changer&amp;range=asc&amp;from=141&amp;to=10" title="Sort the exchangers by the &quot;Exchanger&quot; column" onclick="sort_manual = true; sort_type = 'changer'; sort_range = 'asc'; return update_rates()">Exchanger</a></td>
//<td class="bj from"><a href="https://www.bestchange.com/index.php?sort=from&amp;range=asc&amp;from=141&amp;to=10" title="Sort the exchangers by the &quot;Give&quot; column" onclick="sort_manual = true; sort_type = 'from'; sort_range = 'asc'; return update_rates()">Give</a></td>
//<td class="bj to"><a href="https://www.bestchange.com/index.php?sort=to&amp;range=asc&amp;from=141&amp;to=10" title="Sort the exchangers by the &quot;Get&quot; column" onclick="sort_manual = true; sort_type = 'to'; sort_range = 'asc'; return update_rates()">Get <small>▼</small></a></td>
//<td class="ar arp reserve"><a href="https://www.bestchange.com/index.php?sort=reserve&amp;range=desc&amp;from=141&amp;to=10" title="Sort the exchangers by the &quot;Reserve&quot; column" onclick="sort_manual = true; sort_type = 'reserve'; sort_range = 'desc'; return update_rates()">Reserve</a></td>
//<td class="bj reviews end"><a href="https://www.bestchange.com/index.php?sort=reviews&amp;range=desc&amp;from=141&amp;to=10" title="Sort the exchangers by the &quot;Reviews&quot; column" onclick="sort_manual = true; sort_type = 'reviews'; sort_range = 'desc'; return update_rates()">Reviews</a></td>
//</tr>
//    </thead>
//    <tbody>
//<tr onclick="ccl(871, 141, 10, 95)">
//<td class="ir"><span class="io" id="io0" onmousedown="shc(0)" onclick="stopBubbling(event)" onmouseover="show_info(0)" onmouseout="shd()"></span></td>
//<td class="bj"><div class="pa"><a rel="nofollow" target="_blank" href="https://www.bestchange.com/click.php?id=871&amp;from=141&amp;to=10&amp;city=95" onclick="return fco(871)"></a><div class="pc"><div class="ca" translate="no">BitokMe</div></div></div></td><td class="bi"><div class="fs">1 <small translate="no">EUR in</small> <a class="ct" href="https://www.bestchange.com/euro-cash-to-tether-trc20-in-wars.html" onclick="stopBubbling(event); return true" title="This exchange rate Cash EUR to Tether TRC20 (USDT) available in Warsaw.
//Click to view all exchangers from Warsaw.">Warsaw</a></div><div class="fm"><div class="fm1">min 6 000</div><div class="fm2">max 1 715 317</div></div></td>
//<td class="bi">1.02051230 <small translate="no">USDT TRC20</small></td>
//<td class="ar arp" title="The maximum amount of currency Tether TRC20 (USDT), which can give the BitokMe">1 750 502</td>
//<td class="rw" onclick="crw(0)"><a href="https://www.bestchange.com/bitokme-exchanger.html" class="rwa" onclick="return arw(0)" title="BitokMe reviews"><table><tbody><tr><td class="rwl">0</td><td class="del">/</td><td class="rwr pos">32</td><td class="end"></td></tr></tbody></table></a></td></tr><tr onclick="ccl(850, 141, 10, 54)">
//<td class="ir"><span class="io" id="io1" onmousedown="shc(1)" onclick="stopBubbling(event)" onmouseover="show_info(1)" onmouseout="shd()"></span></td>
//<td class="bj"><div class="pa labpad2"><a rel="nofollow" target="_blank" href="https://www.bestchange.com/click.php?id=850&amp;from=141&amp;to=10&amp;city=54" onclick="return fco(850)"></a><div class="pc"><div class="ca" translate="no">Keine-Exchange</div></div><span class="lbpl" onclick="stopBubbling(event)"><span id="la0" class="verifying" onmouseover="sld(0)" onmouseout="hld()"><span id="ld0">This exchanger can require verification of client's documents<br>or postone the payment for checking other data.</span></span><span id="la1" class="more" onmouseover="sld(1)" onmouseout="hld()"><span id="ld1">This exchanger has <span class="bt">11</span> more exchange rates in other cities.<br>This is the most profitable exchange rate available from the exchanger.</span></span></span></div></td><td class="bi"><div class="fs">1 <small translate="no">EUR in</small> <a class="ct" href="https://www.bestchange.com/euro-cash-to-tether-trc20-in-stam.html" onclick="stopBubbling(event); return true" title="This exchange rate Cash EUR to Tether TRC20 (USDT) available in Istanbul.
//Click to view all exchangers from Istanbul.">Istanbul</a></div><div class="fm"><div class="fm1">min 1 000</div><div class="fm2">max 300 000</div></div></td>
//<td class="bi">1.02000000 <small translate="no">USDT TRC20</small></td>
//<td class="ar arp" title="The maximum amount of currency Tether TRC20 (USDT), which can give the Keine-Exchange">3 433 103</td>
//<td class="rw" onclick="crw(1)"><a href="https://www.bestchange.com/keine-exchange-exchanger.html" class="rwa" onclick="return arw(1)" title="Keine-Exchange reviews"><table><tbody><tr><td class="rwl">0</td><td class="del">/</td><td class="rwr pos">3</td><td class="end"></td></tr></tbody></table></a></td></tr><tr onclick="ccl(877, 141, 10, 199)">
//<td class="ir"><span class="io" id="io2" onmousedown="shc(2)" onclick="stopBubbling(event)" onmouseover="show_info(2)" onmouseout="shd()"></span></td>
//<td class="bj"><div class="pa"><a rel="nofollow" target="_blank" href="https://www.bestchange.com/click.php?id=877&amp;from=141&amp;to=10&amp;city=199" onclick="return fco(877)"></a><div class="pc"><div class="ca" translate="no">Ob-men</div></div></div></td><td class="bi"><div class="fs">1 <small translate="no">EUR in</small> <a class="ct" href="https://www.bestchange.com/euro-cash-to-tether-trc20-in-alan.html" onclick="stopBubbling(event); return true" title="This exchange rate Cash EUR to Tether TRC20 (USDT) available in Alanya.
//Click to view all exchangers from Alanya.">Alanya</a></div><div class="fm"><div class="fm1">min 101</div><div class="fm2">max 8 704</div></div></td>
//<td class="bi">1.01100000 <small translate="no">USDT TRC20</small></td>
//<td class="ar arp" title="The maximum amount of currency Tether TRC20 (USDT), which can give the Ob-men">8 703</td>
//<td class="rw" onclick="crw(2)"><a href="https://www.bestchange.com/ob-men-exchanger.html" class="rwa" onclick="return arw(2)" title="Ob-men reviews"><table><tbody><tr><td class="rwl">0</td><td class="del">/</td><td class="rwr pos">3</td><td class="end"></td></tr></tbody></table></a></td></tr><tr onclick="ccl(1019, 141, 10, 91)">
//<td class="ir"><span class="io" id="io3" onmousedown="shc(3)" onclick="stopBubbling(event)" onmouseover="show_info(3)" onmouseout="shd()"></span></td>
//<td class="bj"><div class="pa labpad1"><a rel="nofollow" target="_blank" href="https://www.bestchange.com/click.php?id=1019&amp;from=141&amp;to=10&amp;city=91" onclick="return fco(1019)"></a><div class="pc"><div class="ca" translate="no">001K</div></div><span class="lbpl" onclick="stopBubbling(event)"><span id="la3" class="more" onmouseover="sld(3)" onmouseout="hld()"><span id="ld3">This exchanger has <span class="bt">18</span> more exchange rates in other cities.<br>This is the most profitable exchange rate available from the exchanger.</span></span></span></div></td><td class="bi"><div class="fs">1 <small translate="no">EUR in</small> <a class="ct" href="https://www.bestchange.com/euro-cash-to-tether-trc20-in-barc.html" onclick="stopBubbling(event); return true" title="This exchange rate Cash EUR to Tether TRC20 (USDT) available in Barcelona.
//Click to view all exchangers from Barcelona.">Barcelona</a></div><div class="fm"><div class="fm1">min 50 000</div><div class="fm2">max 50 000</div></div></td>
//<td class="bi">1 <small translate="no">USDT TRC20</small></td>
//<td class="ar arp" title="The maximum amount of currency Tether TRC20 (USDT), which can give the 001K">551 212</td>
//<td class="rw" onclick="crw(3)"><a href="https://www.bestchange.com/001k-exchanger.html" class="rwa" onclick="return arw(3)" title="001K reviews"><table><tbody><tr><td class="rwl">0</td><td class="del">/</td><td class="rwr pos">1</td><td class="end"></td></tr></tbody></table></a></td></tr><tr onclick="ccl(838, 141, 10, 95)">
//<td class="ir"><span class="io" id="io4" onmousedown="shc(4)" onclick="stopBubbling(event)" onmouseover="show_info(4)" onmouseout="shd()"></span></td>
//<td class="bj"><div class="pa labpad1"><a rel="nofollow" target="_blank" href="https://www.bestchange.com/click.php?id=838&amp;from=141&amp;to=10&amp;city=95" onclick="return fco(838)"></a><div class="pc"><div class="ca" translate="no">AlfaExchange</div></div><span class="lbpl" onclick="stopBubbling(event)"><span id="la4" class="more" onmouseover="sld(4)" onmouseout="hld()"><span id="ld4">This exchanger has <span class="bt">16</span> more exchange rates in other cities.<br>This is the most profitable exchange rate available from the exchanger.</span></span><div id="label_details" class="hide" style="left: 14px; top: -17px;"><table class="stretch_label"><tbody><tr><td class="sl1"></td><td class="sl2"></td><td class="sl3"></td></tr><tr><td class="sl4"></td><td class="sl5"><div id="label_text">This exchanger has <span class="bt">16</span> more exchange rates in other cities.<br>This is the most profitable exchange rate available from the exchanger.</div></td><td class="sl6"></td></tr><tr><td class="sl7"></td><td class="sl8"></td><td class="sl9"></td></tr></tbody></table></div></span></div></td><td class="bi"><div class="fs">1 <small translate="no">EUR in</small> <a class="ct" href="https://www.bestchange.com/euro-cash-to-tether-trc20-in-wars.html" onclick="stopBubbling(event); return true" title="This exchange rate Cash EUR to Tether TRC20 (USDT) available in Warsaw.
//Click to view all exchangers from Warsaw.">Warsaw</a></div><div class="fm"><div class="fm1">min 10 011</div><div class="fm2">max 90 091</div></div></td>
//<td class="bi">0.99899551 <small translate="no">USDT TRC20</small></td>
//<td class="ar arp" title="The maximum amount of currency Tether TRC20 (USDT), which can give the AlfaExchange">11 675 783</td>
//<td class="rw" onclick="crw(4)"><a href="https://www.bestchange.com/alfaexchange-exchanger.html" class="rwa" onclick="return arw(4)" title="AlfaExchange reviews"><table><tbody><tr><td class="rwl">0</td><td class="del">/</td><td class="rwr">0</td><td class="end"></td></tr></tbody></table></a></td></tr><tr onclick="ccl(83, 141, 10, 181)">
//<td class="ir"><span class="io" id="io5" onmousedown="shc(5)" onclick="stopBubbling(event)" onmouseover="show_info(5)" onmouseout="shd()"></span></td>
//<td class="bj"><div class="pa"><a rel="nofollow" target="_blank" href="https://www.bestchange.com/click.php?id=83&amp;from=141&amp;to=10&amp;city=181" onclick="return fco(83)"></a><div class="pc"><div class="ca" translate="no">ObmennikWs</div></div></div></td><td class="bi"><div class="fs">1 <small translate="no">EUR in</small> <a class="ct" href="https://www.bestchange.com/euro-cash-to-tether-trc20-in-wien.html" onclick="stopBubbling(event); return true" title="This exchange rate Cash EUR to Tether TRC20 (USDT) available in Vienna.
//Click to view all exchangers from Vienna.">Vienna</a></div><div class="fm"><div class="fm1">min 10 000</div><div class="fm2">max 300 000</div></div></td>
//<td class="bi">0.99423345 <small translate="no">USDT TRC20</small></td>
//<td class="ar arp" title="The maximum amount of currency Tether TRC20 (USDT), which can give the ObmennikWs">361 230</td>
//<td class="rw" onclick="crw(5)"><a href="https://www.bestchange.com/obmennikws-exchanger.html" class="rwa" onclick="return arw(5)" title="ObmennikWs reviews"><table><tbody><tr><td class="rwl">0</td><td class="del">/</td><td class="rwr pos">1</td><td class="end"></td></tr></tbody></table></a></td></tr><tr onclick="ccl(776, 141, 10, 83)">
//<td class="ir"><span class="io" id="io6" onmousedown="shc(6)" onclick="stopBubbling(event)" onmouseover="show_info(6)" onmouseout="shd()"></span></td>
//<td class="bj"><div class="pa labpad1"><a rel="nofollow" target="_blank" href="https://www.bestchange.com/click.php?id=776&amp;from=141&amp;to=10&amp;city=83" onclick="return fco(776)"></a><div class="pc"><div class="ca" translate="no">XbiTrade</div></div><span class="lbpl" onclick="stopBubbling(event)"><span id="la6" class="more" onmouseover="sld(6)" onmouseout="hld()"><span id="ld6">This exchanger has <span class="bt">8</span> more exchange rates in other cities.<br>This is the most profitable exchange rate available from the exchanger.</span></span></span></div></td><td class="bi"><div class="fs">1 <small translate="no">EUR in</small> <a class="ct" href="https://www.bestchange.com/euro-cash-to-tether-trc20-in-prag.html" onclick="stopBubbling(event); return true" title="This exchange rate Cash EUR to Tether TRC20 (USDT) available in Prague.
//Click to view all exchangers from Prague.">Prague</a></div><div class="fm"><div class="fm1">min 5 059</div><div class="fm2">max 50 586</div></div></td>
//<td class="bi">0.98843000 <small translate="no">USDT TRC20</small></td>
//<td class="ar arp" title="The maximum amount of currency Tether TRC20 (USDT), which can give the XbiTrade">123 802</td>
//<td class="rw" onclick="crw(6)"><a href="https://www.bestchange.com/xbitrade-exchanger.html" class="rwa" onclick="return arw(6)" title="XbiTrade reviews"><table><tbody><tr><td class="rwl">0</td><td class="del">/</td><td class="rwr pos">27</td><td class="end"></td></tr></tbody></table></a></td></tr><tr onclick="ccl(960, 141, 10, 146)">
//<td class="ir"><span class="io" id="io7" onmousedown="shc(7)" onclick="stopBubbling(event)" onmouseover="show_info(7)" onmouseout="shd()"></span></td>
//<td class="bj"><div class="pa labpad1"><a rel="nofollow" target="_blank" href="https://www.bestchange.com/click.php?id=960&amp;from=141&amp;to=10&amp;city=146" onclick="return fco(960)"></a><div class="pc"><div class="ca" translate="no">Crybex</div></div><span class="lbpl" onclick="stopBubbling(event)"><span id="la2" class="more" onmouseover="sld(2)" onmouseout="hld()"><span id="ld2">This exchanger has <span class="bt">2</span> more exchange rates in other cities.<br>This is the most profitable exchange rate available from the exchanger.</span></span></span></div></td><td class="bi"><div class="fs">1 <small translate="no">EUR in</small> <a class="ct" href="https://www.bestchange.com/euro-cash-to-tether-trc20-in-rome.html" onclick="stopBubbling(event); return true" title="This exchange rate Cash EUR to Tether TRC20 (USDT) available in Rome.
//Click to view all exchangers from Rome.">Rome</a></div><div class="fm"><div class="fm1">min 5 000</div><div class="fm2">max 100 000</div></div></td>
//<td class="bi">0.95818390 <small translate="no">USDT TRC20</small></td>
//<td class="ar arp" title="The maximum amount of currency Tether TRC20 (USDT), which can give the Crybex">5 115 707</td>
//<td class="rw" onclick="crw(7)"><a href="https://www.bestchange.com/crybex-exchanger.html" class="rwa" onclick="return arw(7)" title="Crybex reviews"><table><tbody><tr><td class="rwl">0</td><td class="del">/</td><td class="rwr pos">26</td><td class="end"></td></tr></tbody></table></a></td></tr><tr onclick="ccl(341, 141, 10, 2)">
//<td class="ir"><span class="io" id="io8" onmousedown="shc(8)" onclick="stopBubbling(event)" onmouseover="show_info(8)" onmouseout="shd()"></span></td>
//<td class="bj"><div class="pa"><a rel="nofollow" target="_blank" href="https://www.bestchange.com/click.php?id=341&amp;from=141&amp;to=10&amp;city=2" onclick="return fco(341)"></a><div class="pc"><div class="ca" translate="no">WMStream</div></div></div></td><td class="bi"><div class="fs">1 <small translate="no">EUR in</small> <a class="ct" href="https://www.bestchange.com/euro-cash-to-tether-trc20-in-spb.html" onclick="stopBubbling(event); return true" title="This exchange rate Cash EUR to Tether TRC20 (USDT) available in Saint Petersburg.
//Click to view all exchangers from Saint Petersburg.">Saint Petersburg</a></div><div class="fm"><div class="fm1">min 5 000</div><div class="fm2">max 50 000</div></div></td>
//<td class="bi">0.95811868 <small translate="no">USDT TRC20</small></td>
//<td class="ar arp" title="The maximum amount of currency Tether TRC20 (USDT), which can give the WMStream">205 744</td>
//<td class="rw" onclick="crw(8)"><a href="https://www.bestchange.com/wmstream-exchanger.html" class="rwa" onclick="return arw(8)" title="WMStream reviews"><table><tbody><tr><td class="rwl">0</td><td class="del">/</td><td class="rwr">0</td><td class="end"></td></tr></tbody></table></a></td></tr><tr onclick="ccl(1005, 141, 10, 88)">
//<td class="ir"><span class="io" id="io9" onmousedown="shc(9)" onclick="stopBubbling(event)" onmouseover="show_info(9)" onmouseout="shd()"></span></td>
//<td class="bj"><div class="pa"><a rel="nofollow" target="_blank" href="https://www.bestchange.com/click.php?id=1005&amp;from=141&amp;to=10&amp;city=88" onclick="return fco(1005)"></a><div class="pc"><div class="ca" translate="no">EezyCash</div></div></div></td><td class="bi"><div class="fs">1 <small translate="no">EUR in</small> <a class="ct" href="https://www.bestchange.com/euro-cash-to-tether-trc20-in-milan.html" onclick="stopBubbling(event); return true" title="This exchange rate Cash EUR to Tether TRC20 (USDT) available in Milan.
//Click to view all exchangers from Milan.">Milan</a></div><div class="fm"><div class="fm1">min 10 000</div><div class="fm2">max 180 000</div></div></td>
//<td class="bi">0.92729000 <small translate="no">USDT TRC20</small></td>
//<td class="ar arp" title="The maximum amount of currency Tether TRC20 (USDT), which can give the EezyCash">3 000 000</td>
//<td class="rw" onclick="crw(9)"><a href="https://www.bestchange.com/eezycash-exchanger.html" class="rwa" onclick="return arw(9)" title="EezyCash reviews"><table><tbody><tr><td class="rwl">0</td><td class="del">/</td><td class="rwr pos">1</td><td class="end"></td></tr></tbody></table></a></td></tr><tr class="last" onclick="ccl(539, 141, 10, 2)">
//<td class="ir"><span class="io" id="io10" onmousedown="shc(10)" onclick="stopBubbling(event)" onmouseover="show_info(10)" onmouseout="shd()"></span></td>
//<td class="bj"><div class="pa labpad1"><a rel="nofollow" target="_blank" href="https://www.bestchange.com/click.php?id=539&amp;from=141&amp;to=10&amp;city=2" onclick="return fco(539)"></a><div class="pc"><div class="ca" translate="no">NewLine</div></div><span class="lbpl" onclick="stopBubbling(event)"><span id="la5" class="more" onmouseover="sld(5)" onmouseout="hld()"><span id="ld5">This exchanger has <span class="bt">50</span> more exchange rates in other cities.<br>This is the most profitable exchange rate available from the exchanger.</span></span></span></div></td><td class="bi"><div class="fs">1 <small translate="no">EUR in</small> <a class="ct" href="https://www.bestchange.com/euro-cash-to-tether-trc20-in-spb.html" onclick="stopBubbling(event); return true" title="This exchange rate Cash EUR to Tether TRC20 (USDT) available in Saint Petersburg.
//Click to view all exchangers from Saint Petersburg.">Saint Petersburg</a></div><div class="fm"><div class="fm1">min 30 000</div><div class="fm2">max 39 000</div></div></td>
//<td class="bi">0.66995840 <small translate="no">USDT TRC20</small></td>
//<td class="ar arp" title="The maximum amount of currency Tether TRC20 (USDT), which can give the NewLine">3 578 775</td>
//<td class="rw" onclick="crw(10)"><a href="https://www.bestchange.com/newline-exchanger.html" class="rwa" onclick="return arw(10)" title="NewLine reviews"><table><tbody><tr><td class="rwl">0</td><td class="del">/</td><td class="rwr pos">81</td><td class="end"></td></tr></tbody></table></a></td></tr></tbody>
//  </table>
//
//</body>
//</html>`
