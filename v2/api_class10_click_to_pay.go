/*
Decta API

<h2>Introduction</h2><p>This is a documented guide that    describes the process of DECTA card and token operations for both private and business clients. As a new or already    experienced API user you will find in this documentation all necessary and required information to start using DECTA    API and improve your experience with DECTA.</p><h2>Authentication and Authorisation</h2><p>Each request to DECTA API must be signed with a certificate that allows DECTA to identify API clients. This    guarantees the authenticity of the request received from the client.</p><h3>Authentication</h3><p>The authentication process is executed only once for each client.<br>    On successful authentication DECTA assigns a certificate to the client which must be used in every API    request in the future.</p><p>    To access a detailed description of the authentication process, choose <b>Service</b> specification at the top right    corner or click <a        href=\"/?urls.primaryName=SERVICE\" target=\"_blank\">here</a>.</p><h3>Authorisation</h3><p>    Each request must be signed using Jose specifications in particular - <a        href=\"https://tools.ietf.org/html/rfc7515\" target=\"_blank\">RFC7515</a>. It is vital for users to follow the    provided standard and use <b>Base64url    Encoding without Padding</b> for token encoding.</p><h3>Request with filters</h3><p>    Several requests support data filters. Filter syntax can be found below -<p>    <code>        &lt;requestFilter&gt; ::= &lt;filterName&gt;=[eq|ne|gt|ge|lt|le]:&lt;filterValue&gt;    </code></p><ul>Where:    <li>eq - equals</li>    <li>ne - not equals</li>    <li>gt - greater than</li>    <li>ge - greater or equals than</li>    <li>lt - less than</li>    <li>le - less or equals than</li></ul><ul>Example:    <li>cardState: \"eq:BLOCKED_BY_HOLDER\"</li></ul><h2>Changelog</h2>Information about the latest DECTA API changes and improvements can be found <a href=\"/changelog.html\" target=\"_blank\">here.</a>.<h4>Request examples</h4><p>    We recommend users follow the instructions in the next sections below and create the application step by step.</p><p>    To understand the process of DECTA API requests better, we have prepared two basic API calls as examples.</p><p>    <b>POST /v1/api/cards/order</b> request</p><p>For example:</p><ul>    <li>        <b>uri</b> - <br>/v1/api/cards/order    </li>    <li>        <b>body</b> - <br>{\"card\":{\"accountOwnerRelation\":\"OWNER\",\"currencies\":[\"EUR\"],\"deliveryAddress\":{\"city\":\"Madona\",\"country\":\"LVA\",\"name\":\"Test\",\"phone\":9113,\"shipment\":\"STANDARD\",\"street\":\"62        Madona        street\",\"surname\":\"TestS\",\"zipCode\":\"LV-1000\"},\"holder\":{\"currentAddress\":{\"city\":\"Riga\",\"country\":\"LVA\",\"street\":\"7        Street\",\"zipCode\":\"LV0000\"},\"document\":{\"birthDate\":\"1950-12-31\",\"countryCode\":\"EUR\",\"expiryDate\":\"2056-12-31\",\"issuingDate\":\"2016-12-31\",\"number\":\"TTHolder\",\"subtype\":\"RESIDENT_PERMIT_ID\",\"type\":\"DRIVING_LICENSE\"},\"email\":\"oholder@decta.com\",\"language\":\"EN\",\"maidenName\":\"TestM\",\"mobilePhone\":\"12345678\",\"name\":\"TestS\",\"role\":\"OWNER\",\"surname\":\"TestS\",\"userDefinedField1\":\"2\"},\"passphrase\":\"test2pass\",\"priority\":\"URGENT\",\"productCode\":\"701\",\"supplementary\":\"false\"},\"externalId\":\"Private_New\",\"privateClient\":{\"comment\":\"Agreement        num. 123143\",\"currentAddress\":{\"city\":\"RigaCurrent\",\"country\":\"LVA\",\"months\":4,\"street\":\"6 street        street\",\"zipCode\":\"WScur1009\"},\"document\":{\"birthDate\":\"1950-12-31\",\"countryCode\":\"LVA\",\"expiryDate\":\"2099-12-31\",\"issuingDate\":\"1971-12-31\",\"number\":\"1234\",\"subtype\":\"NATIONAL_ID\",\"type\":\"ID_CARD\"},\"email\":\"email1234\",\"language\":\"en\",\"maidenName\":\"Test\",\"mobilePhone\":\"37128866694\",\"name\":\"Test\",\"pep\":\"NO\",\"surname\":\"PrivateBaseNew\",\"userDefinedField1\":\"2000666\"}}    </li></ul><p><h5>Suppose we have:</h5><ul>    <li><b>Private key</b> -<br>308204BE020100300D06092A864886F70D0101010500048204A8308204A40201000282010100C5BD899E1BD0E7D0C7F814976C449288738ECBD396D3BB7ABE8ED6ADACB60AA4622594FB2CE5C97F01274F3F350935A170013743B557882674E2C501C0A17943300C6E39D6D171B06980E12B8946050D4255F6E8E293FCF3DA848DE2486296EB24CA68619ED954804FF00F4936B4289AE04FCA3AD8B0AFA1D3BF7D0DD956D59B528BA284595116B512B5234A53B8CD746268F6F93503925B565DEF8BE9A9A69E0ECE869B40D11F1CF65FC92FDE8159943437F161ED33B4BB83FE5E8F3898E2C03847E84F92BB25CF7835A67F8CA6C24A12F81CC99CE7FBAC387F6D1232668DCA64B1CA0DBFF1292DB7A0C4D129267F6779826AE86337CE00EA789A3F4BCC0E8F0203010001028201001B42BD15320A83DAF1A68ED82D51EACDAAF8B5BAEEE133C8813AD9F7C986B342E451000FFCA99534D028E630F593D673F22808C169DFAE6276DCC8822C57568B5FB56700CC4FC597A783E0070B7EBEB58C3EA3B28982136E5EE2C558BBD7F3AC693F0A52E2FE5D3AEF4A4E212ECC7988988207844359606D33EF0343848207228B91AE9F37D02FD1F5D9994D0DBB9B15DE1F3BDE62980BC3645226896D2AC9A79DE49F0129E2F012D45BEBABA6AEFC7F0E09FB4C551C976BE2AD393EB1D7A448A6579547EA481F87B14EEDC0DB3FB1CAB326F5B4C7B7655BE8D96EE167BA64CFDA033A15F503BED8B62067A162C1ABEEF2B0719360531EB98E127BC21CCF7C1902818100EA311697B86CC3257E2C1BC9CE1928A1ECA2D9F4DC07A5ADA087D0E053AA1F9D3EB46C42627CE5541F386B8ED356033B08D25F195DD9AC27B944FEA16E16C7F8D68FFF43BC76C7B919CA0B6CEAC775E0D59FA55E108182406AFDC3A00C38BA7968A775604813679B0A1BA4389F21B1DDB8002A09FDF94E3875B4D971179645E502818100D8277B9B4BFDB22A2E671BE25BF41D07EBC13D7872C18BBC12C925914594580AB6E0B35BB479F7D437F35F1D5602A4EB04E3B4BA9F3FD66983AFE28D56B9945A55716CE655B6A51926714C97CDA211272A6D8D1A6A5DA903955DDAE4CCF5CB17AB579872ABFB3A4AD2A9B6B7781C217996753D8F08B7FEBACE0437E7E3127B630281804110A25644C6F2F0D7BD297B47E77582875C3771F02FAFB82D818E66C4D7DF30B0DB2FB8C1E43152CDD9BF084F6EF636D09CA20F4A23CFD0B98891B608000C4EF64DCB7DE63E99FEB34B8F211399970680537A7E47D7B37872414EB888ACBA224111456B1A9B498DA9A1445EF6D745B3740B97A26F0D74CDFAE5E403A7A9F88102818100A3196FB867A79F471ACA6C8CE9CB921846F901D41FEB10F031D09B4B39904CBA90F18E04B5728EE10FF2D0A3472229A4B884C4FB486C97F1F80EC048DBD69E2D676975BD7B919C9414A0CB25846134C633CA83FD19A36083B6B2221820DB9A8A5A9C770897E2B2428120D4212E3D33B9162192F899A094520A5FBE7B77A89B7702818100D9DD12BF56479E6875BF4B82450C926F07E56FCC4F8A54594E7C6D76FFBDE50DC1C57097665580D88A9E13989F230EBA30254C056DC664990E06FA54565A977A5693B847F1C3025FF733F24D1F135B51A9FEF2D08476AF26A129F230DFE04140995D0C5BD02FE62067B652822BD25AEDAAAFACE932981FE0FF38A94CE50C0B25    </li>    <li>Signed by Decta certificate <b>SHA256 Thumbprint</b> base64Url encoded -<br>g8Gkc4OqZwIDV7HSeklpQcJZZWyTVkO1JKofXtAN63I    </li></ul><h5>JWS preparing: </h5><ul>    <li>        <b>JWS Header</b> -<br>{\"x5t#S256\":\"g8Gkc4OqZwIDV7HSeklpQcJZZWyTVkO1JKofXtAN63I\",\"alg\":\"RS256\"}    </li>    <li><b>JWS Payload</b>-<br>/v1/api/cards/order{\"card\":{\"accountOwnerRelation\":\"OWNER\",\"currencies\":[\"EUR\"],\"deliveryAddress\":{\"city\":\"Madona\",\"country\":\"LVA\",\"name\":\"Test\",\"phone\":9113,\"shipment\":\"STANDARD\",\"street\":\"62        Madona        street\",\"surname\":\"TestS\",\"zipCode\":\"LV-1000\"},\"holder\":{\"currentAddress\":{\"city\":\"Riga\",\"country\":\"LVA\",\"street\":\"7        Street\",\"zipCode\":\"LV0000\"},\"document\":{\"birthDate\":\"1950-12-31\",\"countryCode\":\"EUR\",\"expiryDate\":\"2056-12-31\",\"issuingDate\":\"2016-12-31\",\"number\":\"TTHolder\",\"subtype\":\"RESIDENT_PERMIT_ID\",\"type\":\"DRIVING_LICENSE\"},\"email\":\"oholder@decta.com\",\"language\":\"EN\",\"maidenName\":\"TestM\",\"mobilePhone\":\"12345678\",\"name\":\"TestS\",\"role\":\"OWNER\",\"surname\":\"TestS\",\"userDefinedField1\":\"2\"},\"passphrase\":\"test2pass\",\"priority\":\"URGENT\",\"productCode\":\"701\",\"supplementary\":\"false\"},\"externalId\":\"Private_New\",\"privateClient\":{\"comment\":\"Agreement        num. 123143\",\"currentAddress\":{\"city\":\"RigaCurrent\",\"country\":\"LVA\",\"months\":4,\"street\":\"6 street        street\",\"zipCode\":\"WScur1009\"},\"document\":{\"birthDate\":\"1950-12-31\",\"countryCode\":\"LVA\",\"expiryDate\":\"2099-12-31\",\"issuingDate\":\"1971-12-31\",\"number\":\"1234\",\"subtype\":\"NATIONAL_ID\",\"type\":\"ID_CARD\"},\"email\":\"email1234\",\"language\":\"en\",\"maidenName\":\"Test\",\"mobilePhone\":\"37128866694\",\"name\":\"Test\",\"pep\":\"NO\",\"surname\":\"PrivateBaseNew\",\"userDefinedField1\":\"2000666\"}}    </li></ul><h5>JWS signing result:</h5><ul>    <li><b>Token-Header</b> (thumbprint encoded using Base64url Encoding without Padding) -<br>eyJ4NXQjUzI1NiI6Imc4R2tjNE9xWndJRFY3SFNla2xwUWNKWlpXeVRWa08xSktvZlh0QU42M0kiLCJhbGciOiJSUzI1NiJ9    </li>    <li style=\"word-wrap: normal\"><b>Token-body</b> (payload encoded using Base64url Encoding without Padding) -<br>L3YxL2FwaS9jYXJkcy9vcmRlcnsiY2FyZCI6eyJhY2NvdW50T3duZXJSZWxhdGlvbiI6Ik9XTkVSIiwiY3VycmVuY2llcyI6WyJFVVIiXSwiZGVsaXZlcnlBZGRyZXNzIjp7ImNpdHkiOiJNYWRvbmEiLCJjb3VudHJ5IjoiTFZBIiwibmFtZSI6IlRlc3QiLCJwaG9uZSI6OTExMywic2hpcG1lbnQiOiJTVEFOREFSRCIsInN0cmVldCI6IjYyIE1hZG9uYSBzdHJlZXQiLCJzdXJuYW1lIjoiVGVzdFMiLCJ6aXBDb2RlIjoiTFYtMTAwMCJ9LCJob2xkZXIiOnsiY3VycmVudEFkZHJlc3MiOnsiY2l0eSI6IlJpZ2EiLCJjb3VudHJ5IjoiTFZBIiwic3RyZWV0IjoiNyBTdHJlZXQiLCJ6aXBDb2RlIjoiTFYwMDAwIn0sImRvY3VtZW50Ijp7ImJpcnRoRGF0ZSI6IjE5NTAtMTItMzEiLCJjb3VudHJ5Q29kZSI6IkVVUiIsImV4cGlyeURhdGUiOiIyMDU2LTEyLTMxIiwiaXNzdWluZ0RhdGUiOiIyMDE2LTEyLTMxIiwibnVtYmVyIjoiVFRIb2xkZXIiLCJzdWJ0eXBlIjoiUkVTSURFTlRfUEVSTUlUX0lEIiwidHlwZSI6IkRSSVZJTkdfTElDRU5TRSJ9LCJlbWFpbCI6Im9ob2xkZXJAZGVjdGEuY29tIiwibGFuZ3VhZ2UiOiJFTiIsIm1haWRlbk5hbWUiOiJUZXN0TSIsIm1vYmlsZVBob25lIjoiMTIzNDU2NzgiLCJuYW1lIjoiVGVzdFMiLCJyb2xlIjoiT1dORVIiLCJzdXJuYW1lIjoiVGVzdFMiLCJ1c2VyRGVmaW5lZEZpZWxkMSI6IjIifSwicGFzc3BocmFzZSI6InRlc3QycGFzcyIsInByaW9yaXR5IjoiVVJHRU5UIiwicHJvZHVjdENvZGUiOiI3MDEiLCJzdXBwbGVtZW50YXJ5IjoiZmFsc2UifSwiZXh0ZXJuYWxJZCI6IlByaXZhdGVfTmV3IiwicHJpdmF0ZUNsaWVudCI6eyJjb21tZW50IjoiQWdyZWVtZW50IG51bS4gMTIzMTQzIiwiY3VycmVudEFkZHJlc3MiOnsiY2l0eSI6IlJpZ2FDdXJyZW50IiwiY291bnRyeSI6IkxWQSIsIm1vbnRocyI6NCwic3RyZWV0IjoiNiBzdHJlZXQgc3RyZWV0IiwiemlwQ29kZSI6IldTY3VyMTAwOSJ9LCJkb2N1bWVudCI6eyJiaXJ0aERhdGUiOiIxOTUwLTEyLTMxIiwiY291bnRyeUNvZGUiOiJMVkEiLCJleHBpcnlEYXRlIjoiMjA5OS0xMi0zMSIsImlzc3VpbmdEYXRlIjoiMTk3MS0xMi0zMSIsIm51bWJlciI6IjEyMzQiLCJzdWJ0eXBlIjoiTkFUSU9OQUxfSUQiLCJ0eXBlIjoiSURfQ0FSRCJ9LCJlbWFpbCI6ImVtYWlsMTIzNCIsImxhbmd1YWdlIjoiZW4iLCJtYWlkZW5OYW1lIjoiVGVzdCIsIm1vYmlsZVBob25lIjoiMzcxMjg4NjY2OTQiLCJuYW1lIjoiVGVzdCIsInBlcCI6Ik5PIiwic3VybmFtZSI6IlByaXZhdGVCYXNlTmV3IiwidXNlckRlZmluZWRGaWVsZDEiOiIyMDAwNjY2In19    </li>    <li><b>token-signature</b> (using encoded token header + encoded JWS Payload separated with a \".\" and signed with        the private key) -<br>wcVPkihLADBs774wPOyqUoK05tj6tSeFWaKuH1jRNED5pABPDpvsg7aDOtbBmORJGF0HfyVfquui5wRKuMQCUW2OaQe26Tlh194lxdfrcg_e1S3P-ZqqI6LEHKLGZFGSuuzT95Da5osSmfMTkr_LOQzT_vPiASr0KFSNSZtL60l2q3dUeZ2AcwTeVFi1Aga_K_mB5UeZUXPQf9XfQ1Eb2aU7Fl9YIsYL2M2B5NmQTgKfy5nuSQXhw3y_ddoiNFR1wHY-g6eX3oU1_kx6s5_GV7Jr9zftj6oCZjhOr-_Ohr2_UMxmpn1D7MGThhwlDguK7CfFWb_8nOvEqntYlTErDQ    </li></ul>Make a request.<p><b>GET /v1/api/cards</b> request<br>    Important - when including query parameters into any of GET requests, they should be preceded by a \"?\" after the URI    in the JWS Payload and combined with \"&\", if multiple are present.<p>For example:</p><ul>    <li>        <b>uri</b> -<br> /v1/api/cards    </li>    <li>        <b>query parameters</b>:<br>        cardAccount: \"eq:00000000\",<br>        cardName: \"eq:Test Test\",<br>        count: \"1000\",<br>        product: \"eq:999\"    </li>    <li>        <b>Resulting URI and query parameters</b>:<br>/v1/api/orders?product=eq:999&cardName=eq:Test        Test&count=1000&cardAccount=eq:00000000    </li></ul><p><h5>Suppose we have:</h5><ul>    <li><b>Private key</b> -<br>308204BE020100300D06092A864886F70D0101010500048204A8308204A40201000282010100C5BD899E1BD0E7D0C7F814976C449288738ECBD396D3BB7ABE8ED6ADACB60AA4622594FB2CE5C97F01274F3F350935A170013743B557882674E2C501C0A17943300C6E39D6D171B06980E12B8946050D4255F6E8E293FCF3DA848DE2486296EB24CA68619ED954804FF00F4936B4289AE04FCA3AD8B0AFA1D3BF7D0DD956D59B528BA284595116B512B5234A53B8CD746268F6F93503925B565DEF8BE9A9A69E0ECE869B40D11F1CF65FC92FDE8159943437F161ED33B4BB83FE5E8F3898E2C03847E84F92BB25CF7835A67F8CA6C24A12F81CC99CE7FBAC387F6D1232668DCA64B1CA0DBFF1292DB7A0C4D129267F6779826AE86337CE00EA789A3F4BCC0E8F0203010001028201001B42BD15320A83DAF1A68ED82D51EACDAAF8B5BAEEE133C8813AD9F7C986B342E451000FFCA99534D028E630F593D673F22808C169DFAE6276DCC8822C57568B5FB56700CC4FC597A783E0070B7EBEB58C3EA3B28982136E5EE2C558BBD7F3AC693F0A52E2FE5D3AEF4A4E212ECC7988988207844359606D33EF0343848207228B91AE9F37D02FD1F5D9994D0DBB9B15DE1F3BDE62980BC3645226896D2AC9A79DE49F0129E2F012D45BEBABA6AEFC7F0E09FB4C551C976BE2AD393EB1D7A448A6579547EA481F87B14EEDC0DB3FB1CAB326F5B4C7B7655BE8D96EE167BA64CFDA033A15F503BED8B62067A162C1ABEEF2B0719360531EB98E127BC21CCF7C1902818100EA311697B86CC3257E2C1BC9CE1928A1ECA2D9F4DC07A5ADA087D0E053AA1F9D3EB46C42627CE5541F386B8ED356033B08D25F195DD9AC27B944FEA16E16C7F8D68FFF43BC76C7B919CA0B6CEAC775E0D59FA55E108182406AFDC3A00C38BA7968A775604813679B0A1BA4389F21B1DDB8002A09FDF94E3875B4D971179645E502818100D8277B9B4BFDB22A2E671BE25BF41D07EBC13D7872C18BBC12C925914594580AB6E0B35BB479F7D437F35F1D5602A4EB04E3B4BA9F3FD66983AFE28D56B9945A55716CE655B6A51926714C97CDA211272A6D8D1A6A5DA903955DDAE4CCF5CB17AB579872ABFB3A4AD2A9B6B7781C217996753D8F08B7FEBACE0437E7E3127B630281804110A25644C6F2F0D7BD297B47E77582875C3771F02FAFB82D818E66C4D7DF30B0DB2FB8C1E43152CDD9BF084F6EF636D09CA20F4A23CFD0B98891B608000C4EF64DCB7DE63E99FEB34B8F211399970680537A7E47D7B37872414EB888ACBA224111456B1A9B498DA9A1445EF6D745B3740B97A26F0D74CDFAE5E403A7A9F88102818100A3196FB867A79F471ACA6C8CE9CB921846F901D41FEB10F031D09B4B39904CBA90F18E04B5728EE10FF2D0A3472229A4B884C4FB486C97F1F80EC048DBD69E2D676975BD7B919C9414A0CB25846134C633CA83FD19A36083B6B2221820DB9A8A5A9C770897E2B2428120D4212E3D33B9162192F899A094520A5FBE7B77A89B7702818100D9DD12BF56479E6875BF4B82450C926F07E56FCC4F8A54594E7C6D76FFBDE50DC1C57097665580D88A9E13989F230EBA30254C056DC664990E06FA54565A977A5693B847F1C3025FF733F24D1F135B51A9FEF2D08476AF26A129F230DFE04140995D0C5BD02FE62067B652822BD25AEDAAAFACE932981FE0FF38A94CE50C0B25    </li>    <li>Signed by Decta certificate <b>SHA256 Thumbprint</b> base64Url encoded -<br>g8Gkc4OqZwIDV7HSeklpQcJZZWyTVkO1JKofXtAN63I    </li></ul><h5>JWS preparing: </h5><ul>    <li>        <b>JWS Header</b> -<br>{\"x5t#S256\":\"g8Gkc4OqZwIDV7HSeklpQcJZZWyTVkO1JKofXtAN63I\",\"alg\":\"RS256\"}    </li>    <li><b>JWS Payload</b> -<br>/v1/api/ordersproduct=eq:999&cardName=eq:Test Test&count=1000&cardAccount=eq:00000000    </li></ul><h5>JWS signing result:</h5><ul>    <li><b>Token-Header</b> (thumbprint encoded using Base64url Encoding without Padding) -<br>eyJ4NXQjUzI1NiI6Imc4R2tjNE9xWndJRFY3SFNla2xwUWNKWlpXeVRWa08xSktvZlh0QU42M0kiLCJhbGciOiJSUzI1NiJ9    </li>    <li style=\"word-wrap: normal\"><b>Token-body</b> (payload encoded using Base64url Encoding without Padding) -<br>L3YxL2FwaS9vcmRlcnNwcm9kdWN0PWVxOjk5OSZjYXJkTmFtZT1lcTpUZXN0IFRlc3QmY291bnQ9MTAwMCZjYXJkQWNjb3VudD1lcTowMDAwMDAwMA    </li>    <li><b>token-signature</b> (using encoded token header + encoded JWS Payload separated with a \".\" and signed with        the private key) -<br>Iw8oJuOZvzxDmCjNZgOSHC-dr_y6hUUTHw8nT8C_bv6aYS9COLOLoTg3ANqo16RrBJM26hrprwMvVjEy1sz2-HiBTnZS6-HQo-KlqvCPqMH7HwSSbwFvedzcjAm2SC1xFUkGnAl5SJ_hxRWAejefmK4X3_ToJx3AN9JbJU40IeRCShKKSrNgui0A_zzkvza-JNiRhzzxRofrNxuUhF_xfGrOiIsGpG3k0l24bjzr-5YRf9vwGhQZRA93_ZAUUoLZ4pKi6CYcBd0dnmTLuIiFIpvBXKt9a-EZb5QJFk8GJnZ7pTXKRuJev2nW76rFhlC6rQB_RVoQTzsaIHae7swycg    </li></ul>Make a request.<p>    <b>If you have further questions regarding DECTA API, please contact our DECTA Support Team -</b><br></p>

API version: 2.16
Contact: support@decta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dectav2

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// Class10ClickToPayAPIService Class10ClickToPayAPI service
type Class10ClickToPayAPIService service

type ApiAddClickToPayRequest struct {
	ctx context.Context
	ApiService *Class10ClickToPayAPIService
	requestId *string
	tokenHeader *string
	tokenSignature *string
	ppan string
	clickToPayCrudApiDto *ClickToPayCrudApiDto
}

// Request ID (UUID format)
func (r ApiAddClickToPayRequest) RequestId(requestId string) ApiAddClickToPayRequest {
	r.requestId = &requestId
	return r
}

// URL64Encoded without padding Header part of JWS token
func (r ApiAddClickToPayRequest) TokenHeader(tokenHeader string) ApiAddClickToPayRequest {
	r.tokenHeader = &tokenHeader
	return r
}

// URL64Encoded without padding Signature part of JWS
func (r ApiAddClickToPayRequest) TokenSignature(tokenSignature string) ApiAddClickToPayRequest {
	r.tokenSignature = &tokenSignature
	return r
}

func (r ApiAddClickToPayRequest) ClickToPayCrudApiDto(clickToPayCrudApiDto ClickToPayCrudApiDto) ApiAddClickToPayRequest {
	r.clickToPayCrudApiDto = &clickToPayCrudApiDto
	return r
}

func (r ApiAddClickToPayRequest) Execute() (*http.Response, error) {
	return r.ApiService.AddClickToPayExecute(r)
}

/*
AddClickToPay Add card to Click to Pay

The request allows users to add a card to Click to Pay.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ppan Masked card number
 @return ApiAddClickToPayRequest
*/
func (a *Class10ClickToPayAPIService) AddClickToPay(ctx context.Context, ppan string) ApiAddClickToPayRequest {
	return ApiAddClickToPayRequest{
		ApiService: a,
		ctx: ctx,
		ppan: ppan,
	}
}

// Execute executes the request
func (a *Class10ClickToPayAPIService) AddClickToPayExecute(r ApiAddClickToPayRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Class10ClickToPayAPIService.AddClickToPay")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/api/click-to-pay/{ppan}"
	localVarPath = strings.Replace(localVarPath, "{"+"ppan"+"}", url.PathEscape(parameterValueToString(r.ppan, "ppan")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.requestId == nil {
		return nil, reportError("requestId is required and must be specified")
	}
	if r.tokenHeader == nil {
		return nil, reportError("tokenHeader is required and must be specified")
	}
	if r.tokenSignature == nil {
		return nil, reportError("tokenSignature is required and must be specified")
	}
	if strlen(r.ppan) < 0 {
		return nil, reportError("ppan must have at least 0 elements")
	}
	if strlen(r.ppan) > 19 {
		return nil, reportError("ppan must have less than 19 elements")
	}
	if r.clickToPayCrudApiDto == nil {
		return nil, reportError("clickToPayCrudApiDto is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Request-Id", r.requestId, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "token-header", r.tokenHeader, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "token-signature", r.tokenSignature, "simple", "")
	// body params
	localVarPostBody = r.clickToPayCrudApiDto
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v UpdateEmail403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v UpdateEmail403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v UpdateEmail403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v UpdateEmail403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiChangeClickToPayRequest struct {
	ctx context.Context
	ApiService *Class10ClickToPayAPIService
	requestId *string
	tokenHeader *string
	tokenSignature *string
	ppan string
	clickToPayCrudApiDto *ClickToPayCrudApiDto
}

// Request ID (UUID format)
func (r ApiChangeClickToPayRequest) RequestId(requestId string) ApiChangeClickToPayRequest {
	r.requestId = &requestId
	return r
}

// URL64Encoded without padding Header part of JWS token
func (r ApiChangeClickToPayRequest) TokenHeader(tokenHeader string) ApiChangeClickToPayRequest {
	r.tokenHeader = &tokenHeader
	return r
}

// URL64Encoded without padding Signature part of JWS
func (r ApiChangeClickToPayRequest) TokenSignature(tokenSignature string) ApiChangeClickToPayRequest {
	r.tokenSignature = &tokenSignature
	return r
}

func (r ApiChangeClickToPayRequest) ClickToPayCrudApiDto(clickToPayCrudApiDto ClickToPayCrudApiDto) ApiChangeClickToPayRequest {
	r.clickToPayCrudApiDto = &clickToPayCrudApiDto
	return r
}

func (r ApiChangeClickToPayRequest) Execute() (*http.Response, error) {
	return r.ApiService.ChangeClickToPayExecute(r)
}

/*
ChangeClickToPay Change Click to Pay information using card

The request allows users to change data saved in Click to Pay.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ppan Masked card number
 @return ApiChangeClickToPayRequest
*/
func (a *Class10ClickToPayAPIService) ChangeClickToPay(ctx context.Context, ppan string) ApiChangeClickToPayRequest {
	return ApiChangeClickToPayRequest{
		ApiService: a,
		ctx: ctx,
		ppan: ppan,
	}
}

// Execute executes the request
func (a *Class10ClickToPayAPIService) ChangeClickToPayExecute(r ApiChangeClickToPayRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Class10ClickToPayAPIService.ChangeClickToPay")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/api/click-to-pay/{ppan}"
	localVarPath = strings.Replace(localVarPath, "{"+"ppan"+"}", url.PathEscape(parameterValueToString(r.ppan, "ppan")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.requestId == nil {
		return nil, reportError("requestId is required and must be specified")
	}
	if r.tokenHeader == nil {
		return nil, reportError("tokenHeader is required and must be specified")
	}
	if r.tokenSignature == nil {
		return nil, reportError("tokenSignature is required and must be specified")
	}
	if strlen(r.ppan) < 0 {
		return nil, reportError("ppan must have at least 0 elements")
	}
	if strlen(r.ppan) > 19 {
		return nil, reportError("ppan must have less than 19 elements")
	}
	if r.clickToPayCrudApiDto == nil {
		return nil, reportError("clickToPayCrudApiDto is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Request-Id", r.requestId, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "token-header", r.tokenHeader, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "token-signature", r.tokenSignature, "simple", "")
	// body params
	localVarPostBody = r.clickToPayCrudApiDto
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v UpdateEmail403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v UpdateEmail403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v UpdateEmail403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v UpdateEmail403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeleteClickToPayRequest struct {
	ctx context.Context
	ApiService *Class10ClickToPayAPIService
	requestId *string
	tokenHeader *string
	tokenSignature *string
	ppan string
	clientId *string
}

// Request ID (UUID format)
func (r ApiDeleteClickToPayRequest) RequestId(requestId string) ApiDeleteClickToPayRequest {
	r.requestId = &requestId
	return r
}

// URL64Encoded without padding Header part of JWS token
func (r ApiDeleteClickToPayRequest) TokenHeader(tokenHeader string) ApiDeleteClickToPayRequest {
	r.tokenHeader = &tokenHeader
	return r
}

// URL64Encoded without padding Signature part of JWS
func (r ApiDeleteClickToPayRequest) TokenSignature(tokenSignature string) ApiDeleteClickToPayRequest {
	r.tokenSignature = &tokenSignature
	return r
}

// clientId is a unique client id number generated on the DECTA partner side. Maximum length 19 symbols
func (r ApiDeleteClickToPayRequest) ClientId(clientId string) ApiDeleteClickToPayRequest {
	r.clientId = &clientId
	return r
}

func (r ApiDeleteClickToPayRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteClickToPayExecute(r)
}

/*
DeleteClickToPay Delete data from Click to Pay using card

The request allows users to delete data saved in Click to Pay.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ppan Masked card number
 @return ApiDeleteClickToPayRequest
*/
func (a *Class10ClickToPayAPIService) DeleteClickToPay(ctx context.Context, ppan string) ApiDeleteClickToPayRequest {
	return ApiDeleteClickToPayRequest{
		ApiService: a,
		ctx: ctx,
		ppan: ppan,
	}
}

// Execute executes the request
func (a *Class10ClickToPayAPIService) DeleteClickToPayExecute(r ApiDeleteClickToPayRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Class10ClickToPayAPIService.DeleteClickToPay")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/api/click-to-pay/{ppan}"
	localVarPath = strings.Replace(localVarPath, "{"+"ppan"+"}", url.PathEscape(parameterValueToString(r.ppan, "ppan")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.requestId == nil {
		return nil, reportError("requestId is required and must be specified")
	}
	if r.tokenHeader == nil {
		return nil, reportError("tokenHeader is required and must be specified")
	}
	if r.tokenSignature == nil {
		return nil, reportError("tokenSignature is required and must be specified")
	}
	if strlen(r.ppan) < 0 {
		return nil, reportError("ppan must have at least 0 elements")
	}
	if strlen(r.ppan) > 19 {
		return nil, reportError("ppan must have less than 19 elements")
	}

	if r.clientId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clientId", r.clientId, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Request-Id", r.requestId, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "token-header", r.tokenHeader, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "token-signature", r.tokenSignature, "simple", "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v UpdateEmail403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v UpdateEmail403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v UpdateEmail403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v UpdateEmail403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetClickToPayRequest struct {
	ctx context.Context
	ApiService *Class10ClickToPayAPIService
	requestId *string
	tokenHeader *string
	tokenSignature *string
	clientId string
}

// Request ID (UUID format)
func (r ApiGetClickToPayRequest) RequestId(requestId string) ApiGetClickToPayRequest {
	r.requestId = &requestId
	return r
}

// URL64Encoded without padding Header part of JWS token
func (r ApiGetClickToPayRequest) TokenHeader(tokenHeader string) ApiGetClickToPayRequest {
	r.tokenHeader = &tokenHeader
	return r
}

// URL64Encoded without padding Signature part of JWS
func (r ApiGetClickToPayRequest) TokenSignature(tokenSignature string) ApiGetClickToPayRequest {
	r.tokenSignature = &tokenSignature
	return r
}

func (r ApiGetClickToPayRequest) Execute() (*ClickToPayGetData, *http.Response, error) {
	return r.ApiService.GetClickToPayExecute(r)
}

/*
GetClickToPay Get Click to Pay information using client id

The request allows users to get data saved in Click to Pay.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param clientId Click to Pay profile id
 @return ApiGetClickToPayRequest
*/
func (a *Class10ClickToPayAPIService) GetClickToPay(ctx context.Context, clientId string) ApiGetClickToPayRequest {
	return ApiGetClickToPayRequest{
		ApiService: a,
		ctx: ctx,
		clientId: clientId,
	}
}

// Execute executes the request
//  @return ClickToPayGetData
func (a *Class10ClickToPayAPIService) GetClickToPayExecute(r ApiGetClickToPayRequest) (*ClickToPayGetData, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ClickToPayGetData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Class10ClickToPayAPIService.GetClickToPay")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/api/click-to-pay/{clientId}"
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterValueToString(r.clientId, "clientId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.requestId == nil {
		return localVarReturnValue, nil, reportError("requestId is required and must be specified")
	}
	if r.tokenHeader == nil {
		return localVarReturnValue, nil, reportError("tokenHeader is required and must be specified")
	}
	if r.tokenSignature == nil {
		return localVarReturnValue, nil, reportError("tokenSignature is required and must be specified")
	}
	if strlen(r.clientId) < 0 {
		return localVarReturnValue, nil, reportError("clientId must have at least 0 elements")
	}
	if strlen(r.clientId) > 19 {
		return localVarReturnValue, nil, reportError("clientId must have less than 19 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Request-Id", r.requestId, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "token-header", r.tokenHeader, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "token-signature", r.tokenSignature, "simple", "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v UpdateEmail403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v UpdateEmail403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v UpdateEmail403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v UpdateEmail403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
