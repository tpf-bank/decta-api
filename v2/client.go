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
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

)

var (
	JsonCheck       = regexp.MustCompile(`(?i:(?:application|text)/(?:[^;]+\+)?json)`)
	XmlCheck        = regexp.MustCompile(`(?i:(?:application|text)/(?:[^;]+\+)?xml)`)
	queryParamSplit = regexp.MustCompile(`(^|&)([^&]+)`)
	queryDescape    = strings.NewReplacer( "%5B", "[", "%5D", "]" )
)

// APIClient manages communication with the Decta API API v2.16
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	Class10ClickToPayAPI *Class10ClickToPayAPIService

	Class1CardOrderAPI *Class1CardOrderAPIService

	Class2CardsAPI *Class2CardsAPIService

	Class3RemindersAPI *Class3RemindersAPIService

	Class4TransactionsAPI *Class4TransactionsAPIService

	Class5AuthorizationsAPI *Class5AuthorizationsAPIService

	Class6ClientsAPI *Class6ClientsAPIService

	Class7TokenAPI *Class7TokenAPIService

	Class8AccountsAPI *Class8AccountsAPIService

	Class9LimitsAPI *Class9LimitsAPIService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.Class10ClickToPayAPI = (*Class10ClickToPayAPIService)(&c.common)
	c.Class1CardOrderAPI = (*Class1CardOrderAPIService)(&c.common)
	c.Class2CardsAPI = (*Class2CardsAPIService)(&c.common)
	c.Class3RemindersAPI = (*Class3RemindersAPIService)(&c.common)
	c.Class4TransactionsAPI = (*Class4TransactionsAPIService)(&c.common)
	c.Class5AuthorizationsAPI = (*Class5AuthorizationsAPIService)(&c.common)
	c.Class6ClientsAPI = (*Class6ClientsAPIService)(&c.common)
	c.Class7TokenAPI = (*Class7TokenAPIService)(&c.common)
	c.Class8AccountsAPI = (*Class8AccountsAPIService)(&c.common)
	c.Class9LimitsAPI = (*Class9LimitsAPIService)(&c.common)

	return c
}

func atoi(in string) (int, error) {
	return strconv.Atoi(in)
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insensitive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.EqualFold(a, needle) {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("expected %s to be of type %s but received %s", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

func parameterValueToString( obj interface{}, key string ) string {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		if actualObj, ok := obj.(interface{ GetActualInstanceValue() interface{} }); ok {
			return fmt.Sprintf("%v", actualObj.GetActualInstanceValue())
		}

		return fmt.Sprintf("%v", obj)
	}
	var param,ok = obj.(MappedNullable)
	if !ok {
		return ""
	}
	dataMap,err := param.ToMap()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%v", dataMap[key])
}

// parameterAddToHeaderOrQuery adds the provided object to the request header or url query
// supporting deep object syntax
func parameterAddToHeaderOrQuery(headerOrQueryParams interface{}, keyPrefix string, obj interface{}, style string, collectionType string) {
	var v = reflect.ValueOf(obj)
	var value = ""
	if v == reflect.ValueOf(nil) {
		value = "null"
	} else {
		switch v.Kind() {
			case reflect.Invalid:
				value = "invalid"

			case reflect.Struct:
				if t,ok := obj.(MappedNullable); ok {
					dataMap,err := t.ToMap()
					if err != nil {
						return
					}
					parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, dataMap, style, collectionType)
					return
				}
				if t, ok := obj.(time.Time); ok {
					parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, t.Format(time.RFC3339Nano), style, collectionType)
					return
				}
				value = v.Type().String() + " value"
			case reflect.Slice:
				var indValue = reflect.ValueOf(obj)
				if indValue == reflect.ValueOf(nil) {
					return
				}
				var lenIndValue = indValue.Len()
				for i:=0;i<lenIndValue;i++ {
					var arrayValue = indValue.Index(i)
					var keyPrefixForCollectionType = keyPrefix
					if style == "deepObject" {
						keyPrefixForCollectionType = keyPrefix + "[" + strconv.Itoa(i) + "]"
					}
					parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefixForCollectionType, arrayValue.Interface(), style, collectionType)
				}
				return

			case reflect.Map:
				var indValue = reflect.ValueOf(obj)
				if indValue == reflect.ValueOf(nil) {
					return
				}
				iter := indValue.MapRange()
				for iter.Next() {
					k,v := iter.Key(), iter.Value()
					parameterAddToHeaderOrQuery(headerOrQueryParams, fmt.Sprintf("%s[%s]", keyPrefix, k.String()), v.Interface(), style, collectionType)
				}
				return

			case reflect.Interface:
				fallthrough
			case reflect.Ptr:
				parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, v.Elem().Interface(), style, collectionType)
				return

			case reflect.Int, reflect.Int8, reflect.Int16,
				reflect.Int32, reflect.Int64:
				value = strconv.FormatInt(v.Int(), 10)
			case reflect.Uint, reflect.Uint8, reflect.Uint16,
				reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				value = strconv.FormatUint(v.Uint(), 10)
			case reflect.Float32, reflect.Float64:
				value = strconv.FormatFloat(v.Float(), 'g', -1, 32)
			case reflect.Bool:
				value = strconv.FormatBool(v.Bool())
			case reflect.String:
				value = v.String()
			default:
				value = v.Type().String() + " value"
		}
	}

	switch valuesMap := headerOrQueryParams.(type) {
		case url.Values:
			if collectionType == "csv" && valuesMap.Get(keyPrefix) != "" {
				valuesMap.Set(keyPrefix, valuesMap.Get(keyPrefix) + "," + value)
			} else {
				valuesMap.Add(keyPrefix, value)
			}
			break
		case map[string]string:
			valuesMap[keyPrefix] = value
			break
	}
}

// helper for converting interface{} parameters to json strings
func parameterToJson(obj interface{}) (string, error) {
	jsonBuf, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(jsonBuf), err
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	if c.cfg.Debug {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		log.Printf("\n%s\n", string(dump))
	}

	resp, err := c.cfg.HTTPClient.Do(request)
	if err != nil {
		return resp, err
	}

	if c.cfg.Debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return resp, err
		}
		log.Printf("\n%s\n", string(dump))
	}
	return resp, err
}

// Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
func (c *APIClient) GetConfig() *Configuration {
	return c.cfg
}

type formFile struct {
		fileBytes []byte
		fileName string
		formFileName string
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	formFiles []formFile) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// add form parameters and file if available.
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(formParams) > 0 || (len(formFiles) > 0) {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and multipart form at the same time.")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					w.WriteField(k, iv)
				}
			}
		}
		for _, formFile := range formFiles {
			if len(formFile.fileBytes) > 0 && formFile.fileName != "" {
				w.Boundary()
				part, err := w.CreateFormFile(formFile.formFileName, filepath.Base(formFile.fileName))
				if err != nil {
						return nil, err
				}
				_, err = part.Write(formFile.fileBytes)
				if err != nil {
						return nil, err
				}
			}
		}

		// Set the Boundary in the Content-Type
		headerParams["Content-Type"] = w.FormDataContentType()

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		w.Close()
	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and x-www-form-urlencoded form at the same time.")
		}
		body = &bytes.Buffer{}
		body.WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		url.Host = c.cfg.Host
	}

	// Override request scheme, if applicable
	if c.cfg.Scheme != "" {
		url.Scheme = c.cfg.Scheme
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = queryParamSplit.ReplaceAllStringFunc(query.Encode(), func(s string) string {
		pieces := strings.Split(s, "=")
		pieces[0] = queryDescape.Replace(pieces[0])
		return strings.Join(pieces, "=")
	})

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers[h] = []string{v}
		}
		localVarRequest.Header = headers
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", c.cfg.UserAgent)

	if ctx != nil {
		// add context to the request
		localVarRequest = localVarRequest.WithContext(ctx)

		// Walk through any authentication.

	}

	for header, value := range c.cfg.DefaultHeader {
		localVarRequest.Header.Add(header, value)
	}
	return localVarRequest, nil
}

func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
	if len(b) == 0 {
		return nil
	}
	if s, ok := v.(*string); ok {
		*s = string(b)
		return nil
	}
	if f, ok := v.(*os.File); ok {
		f, err = os.CreateTemp("", "HttpClientFile")
		if err != nil {
			return
		}
		_, err = f.Write(b)
		if err != nil {
			return
		}
		_, err = f.Seek(0, io.SeekStart)
		return
	}
	if f, ok := v.(**os.File); ok {
		*f, err = os.CreateTemp("", "HttpClientFile")
		if err != nil {
			return
		}
		_, err = (*f).Write(b)
		if err != nil {
			return
		}
		_, err = (*f).Seek(0, io.SeekStart)
		return
	}
	if XmlCheck.MatchString(contentType) {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}
	if JsonCheck.MatchString(contentType) {
		if actualObj, ok := v.(interface{ GetActualInstance() interface{} }); ok { // oneOf, anyOf schemas
			if unmarshalObj, ok := actualObj.(interface{ UnmarshalJSON([]byte) error }); ok { // make sure it has UnmarshalJSON defined
				if err = unmarshalObj.UnmarshalJSON(b); err != nil {
					return err
				}
			} else {
				return errors.New("Unknown type with GetActualInstance but no unmarshalObj.UnmarshalJSON defined")
			}
		} else if err = json.Unmarshal(b, v); err != nil { // simple model
			return err
		}
		return nil
	}
	return errors.New("undefined response type")
}

// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if fp, ok := body.(*os.File); ok {
		_, err = bodyBuf.ReadFrom(fp)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if JsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if XmlCheck.MatchString(contentType) {
		var bs []byte
		bs, err = xml.Marshal(body)
		if err == nil {
			bodyBuf.Write(bs)
		}
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("invalid body type %s\n", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		} else {
			expires = now.Add(lifetime)
		}
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}

// GenericOpenAPIError Provides access to the body, error and model on returned errors.
type GenericOpenAPIError struct {
	body  []byte
	error string
	model interface{}
}

// Error returns non-empty string if there was an error.
func (e GenericOpenAPIError) Error() string {
	return e.error
}

// Body returns the raw bytes of the response
func (e GenericOpenAPIError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericOpenAPIError) Model() interface{} {
	return e.model
}

// format error message using title and detail when model implements rfc7807
func formatErrorMessage(status string, v interface{}) string {
	str := ""
	metaValue := reflect.ValueOf(v).Elem()

	if metaValue.Kind() == reflect.Struct {
		field := metaValue.FieldByName("Title")
		if field != (reflect.Value{}) {
			str = fmt.Sprintf("%s", field.Interface())
		}

		field = metaValue.FieldByName("Detail")
		if field != (reflect.Value{}) {
			str = fmt.Sprintf("%s (%s)", str, field.Interface())
		}
	}

	return strings.TrimSpace(fmt.Sprintf("%s %s", status, str))
}
